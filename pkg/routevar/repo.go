package routevar

import (
	"path"
	"regexp"
	"strings"
)

// A RepoRev specifies a repo at a revision. Unlike
// sourcegraph.RepoRevSpec, the revision need not be an absolute
// commit ID. This RepoRev type is appropriate for user input (e.g.,
// from a URL), where it is convenient to allow users to specify
// non-absolute commit IDs that the server can resolve.
type RepoRev struct {
	Repo string // a repo path
	Rev  string // a VCS revision specifier (branch, "master~7", commit ID, etc.)
}

// A TreeEntry specifies a tree/blob by path in a repo at a
// revision. It is analogous to sourcegraph.TreeEntrySpec, but like
// RepoRev, it allows for a non-absolute commit ID.
type TreeEntry struct {
	RepoRev
	Path string // path to the VCS tree/blob
}

var (
	Repo = `{Repo:` + namedToNonCapturingGroups(RepoPattern) + `}`
	Rev  = `{Rev:` + namedToNonCapturingGroups(RevPattern) + `}`

	RepoRevSuffix = `{Rev:` + namedToNonCapturingGroups(`(?:@`+RevPattern+`)?`) + `}`
)

const (
	// RepoPattern is the regexp pattern that matches RepoSpec strings
	// ("repo" or "domain.com/repo" or "domain.com/path/to/repo").
	RepoPattern = `(?P<repo>(?:` + pathComponentNotDelim + `/)*` + pathComponentNotDelim + `)`

	RepoPathDelim         = "-"
	pathComponentNotDelim = `(?:[^@/` + RepoPathDelim + `]|(?:[^/@]{2,}))`

	// RevPattern is the regexp pattern that matches a VCS revision
	// specifier (e.g., "master" or "my/branch~1", or a full 40-char
	// commit ID).
	RevPattern = `(?P<rev>(?:` + pathComponentNotDelim + `/)*` + pathComponentNotDelim + `)`

	// CommitPattern is the regexp pattern that matches absolute
	// (40-character) hexidecimal commit IDs.
	CommitPattern = `(?P<commit>[[:xdigit:]]{40})`
)

var (
	repoPattern = regexp.MustCompile("^" + RepoPattern + "$")
	revPattern  = regexp.MustCompile("^" + RevPattern + "$")
)

// ParseRepo parses a RepoSpec string. If spec is invalid, an
// InvalidError is returned.
func ParseRepo(spec string) (repo string, err error) {
	if m := repoPattern.FindStringSubmatch(spec); len(m) > 0 {
		repo = m[0]
		return
	}
	return "", InvalidError{"RepoSpec", spec, nil}
}

// RepoRouteVars returns route variables for constructing repository
// routes.
func RepoRouteVars(repo string) map[string]string {
	return map[string]string{"Repo": repo}
}

// ToRepoRev marshals a map containing route variables
// generated by (RepoRevSpec).RouteVars() and returns the equivalent
// RepoRevSpec struct.
func ToRepoRev(routeVars map[string]string) RepoRev {
	rr := RepoRev{Repo: ToRepo(routeVars)}
	if revStr := routeVars["Rev"]; revStr != "" {
		if !strings.HasPrefix(revStr, "@") {
			panic("Rev should have had '@' prefix from route")
		}
		rr.Rev = strings.TrimPrefix(revStr, "@")
	}
	if _, ok := routeVars["CommitID"]; ok {
		panic("unexpected CommitID route var; was removed in the simple-routes branch")
	}
	return rr
}

// ToRepo returns the repo path string from a map containing route variables.
func ToRepo(routeVars map[string]string) string {
	return routeVars["Repo"]
}

// RepoRevRouteVars returns route variables for constructing routes to a
// repository revision.
func RepoRevRouteVars(s RepoRev) map[string]string {
	m := RepoRouteVars(s.Repo)
	var rev string
	if s.Rev != "" {
		rev = "@" + s.Rev
	}
	m["Rev"] = rev
	return m
}

func ToTreeEntry(routeVars map[string]string) TreeEntry {
	rr := ToRepoRev(routeVars)
	return TreeEntry{
		RepoRev: rr,
		Path:    path.Clean(strings.TrimPrefix(routeVars["Path"], "/")),
	}
}

func TreeEntryRouteVars(s TreeEntry) map[string]string {
	m := RepoRevRouteVars(s.RepoRev)
	m["Path"] = s.Path
	return m
}
