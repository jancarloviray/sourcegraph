package httpapi

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"

	"sourcegraph.com/sourcegraph/sourcegraph/api/sourcegraph"
	"sourcegraph.com/sourcegraph/srclib/graph"
	srclibstore "sourcegraph.com/sourcegraph/srclib/store"
	"sourcegraph.com/sourcegraph/srclib/store/pb"
	"sourcegraph.com/sourcegraph/srclib/unit"
	"sourcegraph.com/sqs/pbtypes"
)

func TestSrclibImport(t *testing.T) {
	c, mock := newTest()

	const (
		wantRepo     = "r"
		wantCommitID = "c"
	)

	// Sample srclib data to import.
	files := map[string]interface{}{
		"a/b.unit.json":    &unit.SourceUnit{Key: unit.Key{Name: "a/b", Type: "t"}, Info: unit.Info{Dir: ".", Files: []string{"f"}}},
		"a/b/t.graph.json": graph.Output{Defs: []*graph.Def{{DefKey: graph.DefKey{Path: "p"}, Name: "n", File: "f"}}},
	}

	calledReposResolve := mock.Repos.MockResolve_Local(t, wantRepo, 1)
	calledReposGet := mock.Repos.MockGet_Return(t, &sourcegraph.Repo{
		ID:   1,
		URI:  wantRepo,
		Fork: true, // specify fork to prevent background index refreshes
	})
	calledReposResolveRev := mock.Repos.MockResolveRev_NoCheck(t, wantCommitID)
	mock.Search.RefreshIndex_ = func(ctx context.Context, in *sourcegraph.SearchRefreshIndexOp) (*pbtypes.Void, error) {
		return nil, nil
	}

	// Mock the srclib store interface (and replace the old
	// newSrclibStoreClient value when done).
	var calledSrclibStoreImport, calledSrclibStoreIndex, calledSrclibStoreCreateVersion bool
	orig := newSrclibStoreClient
	newSrclibStoreClient = func(context.Context, pb.MultiRepoImporterClient) pb.MultiRepoImporterIndexer {
		return srclibstore.MockMultiRepoStore{
			Import_: func(repo, commitID string, unit *unit.SourceUnit, data graph.Output) error {
				calledSrclibStoreImport = true
				if repo != wantRepo {
					t.Errorf("got repo %q, want %q", repo, wantRepo)
				}
				if commitID != wantCommitID {
					t.Errorf("got commitID %q, want %q", commitID, wantCommitID)
				}
				if want := files["a/b.unit.json"]; !deepEqual(unit, want) {
					t.Errorf("got unit %+v, want %+v", unit, want)
				}
				if want := files["a/b/t.graph.json"]; !deepEqual(data, want) {
					t.Errorf("got graph data %+v, want %+v", data, want)
				}
				return nil
			},
			Index_: func(repo, commitID string) error {
				calledSrclibStoreIndex = true
				if repo != wantRepo {
					t.Errorf("got repo %q, want %q", repo, wantRepo)
				}
				if commitID != wantCommitID {
					t.Errorf("got commitID %q, want %q", commitID, wantCommitID)
				}
				return nil
			},
			CreateVersion_: func(repo, commitID string) error {
				calledSrclibStoreCreateVersion = true
				if repo != wantRepo {
					t.Errorf("got repo %q, want %q", repo, wantRepo)
				}
				if commitID != wantCommitID {
					t.Errorf("got commitID %q, want %q", commitID, wantCommitID)
				}
				return nil
			},
		}
	}
	defer func() { newSrclibStoreClient = orig }()

	// Create a dummy srclib archive.
	var zipData bytes.Buffer
	zipW := zip.NewWriter(&zipData)
	for name, v := range files {
		f, err := zipW.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.NewEncoder(f).Encode(v); err != nil {
			t.Fatal(err)
		}
	}
	if err := zipW.Close(); err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest("PUT", "/repos/r@v/-/srclib-import", &zipData)
	req.Header.Set("content-type", "application/x-zip-compressed")
	req.Header.Set("content-transfer-encoding", "binary")
	if _, err := c.DoOK(req); err != nil {
		t.Fatal(err)
	}
	if !calledSrclibStoreImport {
		t.Error("!calledSrclibStoreImport")
	}
	if !calledSrclibStoreIndex {
		t.Error("!calledSrclibStoreIndex")
	}
	if !calledSrclibStoreCreateVersion {
		t.Error("!calledSrclibStoreCreateVersion")
	}
	if !*calledReposResolve {
		t.Error("!calledReposResolve")
	}
	if !*calledReposGet {
		t.Error("!calledReposGet")
	}
	if !*calledReposResolveRev {
		t.Error("!calledReposResolveRev")
	}
}

func TestSrclibImport_empty(t *testing.T) {
	c, mock := newTest()

	calledReposResolve := mock.Repos.MockResolve_Local(t, "r", 1)
	calledReposGet := mock.Repos.MockGet(t, 1)
	calledReposResolveRev := mock.Repos.MockResolveRev_NoCheck(t, "c")

	// POST an empty zip archive.
	req, _ := http.NewRequest("PUT", "/repos/r@v/-/srclib-import", nil)
	req.Header.Set("content-type", "application/x-zip-compressed")
	req.Header.Set("content-transfer-encoding", "binary")
	resp, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if want := http.StatusBadRequest; resp.StatusCode != want {
		t.Errorf("got HTTP response status %d, want %d", resp.StatusCode, want)
	}
	if !*calledReposResolve {
		t.Error("!calledReposResolve")
	}
	if !*calledReposGet {
		t.Error("!calledReposGet")
	}
	if !*calledReposResolveRev {
		t.Error("!calledReposResolveRev")
	}
}

func deepEqual(u, v interface{}) bool {
	u_, err := json.Marshal(u)
	if err != nil {
		return false
	}
	v_, err := json.Marshal(v)
	if err != nil {
		return false
	}
	return string(u_) == string(v_)
}
