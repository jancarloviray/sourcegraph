package handlerutil

import (
	"bytes"
	"fmt"
	"go/doc"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"gopkg.in/inconshreveable/log15.v2"
	"sourcegraph.com/sourcegraph/sourcegraph/api/sourcegraph"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/errcode"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/htmlutil"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/router_util"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/routevar"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/vcs"
	"sourcegraph.com/sourcegraph/srclib/graph"
)

// RepoCommon holds all of the information necessary to render a
// repository page template. It is returned by GetRepoFromRequest. See also
// RepoRevCommon.
type RepoCommon struct {
	Repo       *sourcegraph.Repo
	RepoConfig *sourcegraph.RepoConfig
}

// RepoRevCommon holds all of the commit-specific information
// necessary to render a repository page template for a certain
// commit. It is returned by GetRepoAndRevFromRequest. It is assumed that pages
// rendered are also provided with repoCommon template data.
type RepoRevCommon struct {
	RepoRevSpec sourcegraph.RepoRevSpec
}

// GetRepoAndRevCommon returns the repository and RepoRevSpec based on
// the route vars. It may also return custom error types
// URLMovedError, NoVCSDataError, which callers should ideally check
// for.
func GetRepoAndRevCommon(ctx context.Context, vars map[string]string) (rc *RepoCommon, vc *RepoRevCommon, err error) {
	rc, err = GetRepoCommon(ctx, vars)
	if err != nil {
		return
	}

	vc = &RepoRevCommon{}
	vc.RepoRevSpec.RepoSpec = rc.Repo.RepoSpec()

	vc.RepoRevSpec, err = getRepoRev(ctx, vars, rc.Repo.DefaultBranch)
	if err != nil {
		cloneInProgress := grpc.Code(err) == codes.Unavailable && grpc.ErrorDesc(err) == vcs.RepoNotExistError{CloneInProgress: true}.Error()
		if noVCSData := grpc.Code(err) == codes.NotFound ||
			cloneInProgress ||
			strings.Contains(err.Error(), "has no default branch"); noVCSData {

			if cloneInProgress {
				return rc, vc, err
			} else if rev := vars["Rev"]; rev != "" && rev != "@" {
				err = vcs.ErrRevisionNotFound
			} else {
				err = &NoVCSDataError{RepoCommon: rc}
			}
		}
		return
	}

	return
}

// GetRepoCommon returns the repository and RepoSpec based on the
// route vars. Callers should ideally handle the custom error type
// URLMovedError.
func GetRepoCommon(ctx context.Context, vars map[string]string) (rc *RepoCommon, err error) {
	cl, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}

	rc = &RepoCommon{}
	rc.Repo, _, err = GetRepo(ctx, vars)
	if err != nil {
		return
	}

	repoSpec := rc.Repo.RepoSpec()
	rc.RepoConfig, err = cl.Repos.GetConfig(ctx, &repoSpec)
	return
}

// GetRepo gets the repo (from the reposSvc) specified in the URL's
// RepoSpec route param. Callers should ideally check for a return error of type
// URLMovedError and handle this scenario by warning or redirecting the user.
func GetRepo(ctx context.Context, vars map[string]string) (repo *sourcegraph.Repo, repoPath string, err error) {
	origRepoPath := routevar.ToRepo(vars)
	cl, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, "", err
	}

	repoPath = origRepoPath
	repo, err = cl.Repos.Get(ctx, &sourcegraph.RepoSpec{URI: repoPath})
	if err != nil {
		return nil, origRepoPath, err
	}
	repoPath = repo.URI

	// Check for redirect.
	if origRepoPath != "" && origRepoPath != repoPath {
		return nil, repoPath, &URLMovedError{repoPath}
	}

	return repo, repoPath, nil
}

// getRepoRev resolves the RepoRevSpec and commit specified in the
// route vars. The provided defaultBranch is used if no rev is
// specified in the URL.
func getRepoRev(ctx context.Context, vars map[string]string, defaultRev string) (sourcegraph.RepoRevSpec, error) {
	repoRev := routevar.ToRepoRev(vars)
	if repoRev.Rev == "" {
		repoRev.Rev = defaultRev

		if repoRev.Rev == "" {
			log15.Warn("getRepoRev: no rev specified and repo has no default rev", "repo", repoRev.Repo)
		}
	}

	cl, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return sourcegraph.RepoRevSpec{}, err
	}

	res, err := cl.Repos.ResolveRev(ctx, &sourcegraph.ReposResolveRevOp{
		Repo: sourcegraph.RepoSpec{URI: repoRev.Repo},
		Rev:  repoRev.Rev,
	})
	if err != nil {
		return sourcegraph.RepoRevSpec{}, err
	}

	return sourcegraph.RepoRevSpec{RepoSpec: sourcegraph.RepoSpec{URI: repoRev.Repo}, CommitID: res.CommitID}, nil
}

// GetRepoAndRev returns the Repo and the RepoRevSpec for a repository. It may
// also return custom error URLMovedError to allow special handling of this case,
// such as for example redirecting the user.
func GetRepoAndRev(ctx context.Context, vars map[string]string) (repo *sourcegraph.Repo, repoRevSpec sourcegraph.RepoRevSpec, err error) {
	var repoPath string
	repo, repoPath, err = GetRepo(ctx, vars)
	if err != nil {
		return repo, repoRevSpec, err
	}
	repoRevSpec.URI = repoPath

	repoRevSpec, err = getRepoRev(ctx, vars, repo.DefaultBranch)
	return repo, repoRevSpec, err
}

// RedirectToNewRepoURI writes an HTTP redirect response with a
// Location that matches the request's location except with the
// RepoSpec route var updated to refer to newRepoURI (instead of the
// originally requested repo URI).
func RedirectToNewRepoURI(w http.ResponseWriter, r *http.Request, newRepoURI string) error {
	origVars := mux.Vars(r)
	origVars["Repo"] = newRepoURI

	destURL, err := mux.CurrentRoute(r).URLPath(router_util.MapToArray(origVars)...)
	if err != nil {
		return err
	}

	http.Redirect(w, r, destURL.String(), http.StatusMovedPermanently)
	return nil
}

// ResolveSrclibDataVersion calls Repos.GetSrclibDataVersionForPath on
// the given entry spec. If a srclib data version exists,
// entry.RepoRev.CommitID is set to the version's commit ID.
//
// If the rev requested by the user (userInputRev) is empty, then it
// performs a lenient resolution: it behaves as though entry.Path is
// also empty and returns the latest build on the default branch for
// the repository, even if the requested file has changed more
// recently. This is because in this case we assume the user cares
// more about seeing srclib defs/refs than any exact version.
func ResolveSrclibDataVersion(ctx context.Context, entry sourcegraph.TreeEntrySpec, userInputRev string) (sourcegraph.RepoRevSpec, *sourcegraph.SrclibDataVersion, error) {
	cl, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return sourcegraph.RepoRevSpec{}, nil, err
	}

	if userInputRev == "" {
		entry.Path = ""
	}

	dataVer, err := cl.Repos.GetSrclibDataVersionForPath(ctx, &entry)
	if err == nil {
		entry.RepoRev.CommitID = dataVer.CommitID
	}
	return entry.RepoRev, dataVer, err
}

// GetDefCommon returns common information about a definition, based
// on the route vars.  It additionally returns common repository and
// revision information. It may also return custom errors
// URLMovedError, or NoVCSDataError.
//
// dc.Def.DefKey will be set to the def specification based on the
// request when getting actual def fails.
func GetDefCommon(ctx context.Context, vars map[string]string, opt *sourcegraph.DefGetOptions) (dc *sourcegraph.Def, err error) {
	repoRev := routevar.ToRepoRev(vars)
	defSpec := routevar.ToDefAtRev(vars)

	// If we fail to get a def, return the best known information to the caller.
	dc = &sourcegraph.Def{
		Def: graph.Def{
			DefKey: graph.DefKey{
				Repo:     defSpec.Repo,
				Unit:     defSpec.Unit,
				UnitType: defSpec.UnitType,
				Path:     defSpec.Path,
			},
		},
	}

	cl, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return
	}

	res, err := cl.Repos.ResolveRev(ctx, &sourcegraph.ReposResolveRevOp{
		Repo: sourcegraph.RepoSpec{URI: repoRev.Repo},
		Rev:  repoRev.Rev,
	})
	if err != nil {
		return
	}
	absRepoRev := sourcegraph.RepoRevSpec{
		RepoSpec: sourcegraph.RepoSpec{URI: repoRev.Repo},
		CommitID: res.CommitID,
	}

	resolvedRev, _, err := ResolveSrclibDataVersion(ctx, sourcegraph.TreeEntrySpec{RepoRev: absRepoRev}, repoRev.Rev)
	if err != nil {
		return dc, err
	}
	dc.Def.DefKey.CommitID = resolvedRev.CommitID

	dc, err = cl.Defs.Get(ctx, &sourcegraph.DefsGetOp{
		Def: sourcegraph.NewDefSpecFromDefKey(dc.Def.DefKey),
		Opt: opt,
	})
	if err != nil {
		return dc, err
	}

	if repoRev.Rev != "" {
		// Check if the def's file has been changed AFTER the resolved
		// srclib-last-version AND the user specified a rev in the URL. If
		// so, then we can't actually display this def, because we'd only
		// be able to show it on an older version of the file (which would
		// mean that users would see file data from an older commit when
		// looking at a newer commit's def--that is BAD).
		//
		// Right now, the best course of action is to 404. This is a
		// fairly rare case that should be remedied as soon as the next
		// build completes. The alternative would be to display a warning
		// saying "this file is N commits behind the requested commit,"
		// but that adds a lot of complexity to the code and to the UI (as
		// we have seen in the past). If a user's looking at a file that
		// was changed since the last srclib version, they also see an
		// unannotated file, so this is consistent with that behavior as
		// well.
		//
		// If the user didn't specify a rev in the URL, then they probably
		// care more about seeing the def than seeing the exact version,
		// so we only perform this strict check if they did.
		defResolvedRev, err := cl.Repos.GetSrclibDataVersionForPath(ctx, &sourcegraph.TreeEntrySpec{
			RepoRev: absRepoRev, // use originally requested rev, not already resolved last-srclib-version
			Path:    dc.File,
		})
		if err != nil {
			return dc, err
		}
		if defResolvedRev.CommitID != resolvedRev.CommitID {
			return dc, &errcode.HTTPErr{
				Status: http.StatusNotFound,
				Err:    fmt.Errorf("no srclib data for def %v (file %s was modified between last srclib analysis version %s and rev %s)", defSpec, dc.File, resolvedRev.CommitID, repoRev.Rev),
			}
		}
	}

	// this can not be moved to svc/local, because HTML sanitation needs to
	// happen on the local sourcegraph instance, not on an untrusted
	// server
	if len(dc.Docs) > 0 {
		defDoc := dc.Docs[0]
		var docHTML string
		switch defDoc.Format {
		case "text/html":
			docHTML = defDoc.Data
		// TODO "text/x-markdown"
		// TODO "text/x-rst"
		default: // including "text/plain"
			var buf bytes.Buffer
			doc.ToHTML(&buf, defDoc.Data, nil)
			docHTML = buf.String()
		}
		dc.DocHTML = htmlutil.SanitizeForPB(docHTML)
	}
	return dc, nil
}
