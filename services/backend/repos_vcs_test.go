package backend

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"sourcegraph.com/sourcegraph/sourcegraph/api/sourcegraph"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/vcs"
	vcstest "sourcegraph.com/sourcegraph/sourcegraph/pkg/vcs/testing"
)

func TestReposService_resolveRev_noRevSpecified_getsDefaultBranch(t *testing.T) {
	ctx, mock := testContext()

	want := strings.Repeat("a", 40)

	calledGet := mock.servers.Repos.MockGet_Return(t, &sourcegraph.Repo{ID: 1, URI: "r", DefaultBranch: "b"})
	var calledVCSRepoResolveRevision bool
	mock.stores.RepoVCS.MockOpen(t, 1, vcstest.MockRepository{
		ResolveRevision_: func(rev string) (vcs.CommitID, error) {
			calledVCSRepoResolveRevision = true
			return vcs.CommitID(want), nil
		},
	})

	// (no rev/branch specified)
	commitID, err := resolveRepoRev(ctx, 1, "")
	if err != nil {
		t.Fatal(err)
	}
	if !*calledGet {
		t.Error("!calledGet")
	}
	if !calledVCSRepoResolveRevision {
		t.Error("!calledVCSRepoResolveRevision")
	}
	if string(commitID) != want {
		t.Errorf("got resolved commit %q, want %q", commitID, want)
	}
}

func TestReposService_resolveRev_noCommitIDSpecified_resolvesRev(t *testing.T) {
	ctx, mock := testContext()

	want := strings.Repeat("a", 40)

	calledGet := mock.servers.Repos.MockGet(t, 1)
	var calledVCSRepoResolveRevision bool
	mock.stores.RepoVCS.MockOpen(t, 1, vcstest.MockRepository{
		ResolveRevision_: func(rev string) (vcs.CommitID, error) {
			calledVCSRepoResolveRevision = true
			return vcs.CommitID(want), nil
		},
	})

	commitID, err := resolveRepoRev(ctx, 1, "b")
	if err != nil {
		t.Fatal(err)
	}
	if !*calledGet {
		t.Error("!calledGet")
	}
	if !calledVCSRepoResolveRevision {
		t.Error("!calledVCSRepoResolveRevision")
	}
	if string(commitID) != want {
		t.Errorf("got resolved commit %q, want %q", commitID, want)
	}
}

func TestReposService_resolveRev_commitIDSpecified_resolvesCommitID(t *testing.T) {
	ctx, mock := testContext()

	want := strings.Repeat("a", 40)

	calledGet := mock.servers.Repos.MockGet(t, 1)
	var calledVCSRepoResolveRevision bool
	mock.stores.RepoVCS.MockOpen(t, 1, vcstest.MockRepository{
		ResolveRevision_: func(rev string) (vcs.CommitID, error) {
			calledVCSRepoResolveRevision = true
			return vcs.CommitID(want), nil
		},
	})

	commitID, err := resolveRepoRev(ctx, 1, strings.Repeat("a", 40))
	if err != nil {
		t.Fatal(err)
	}
	if !*calledGet {
		t.Error("!calledGet")
	}
	if !calledVCSRepoResolveRevision {
		t.Error("!calledVCSRepoResolveRevision")
	}
	if string(commitID) != want {
		t.Errorf("got resolved commit %q, want %q", commitID, want)
	}
}

func TestReposService_resolveRev_commitIDSpecified_failsToResolve(t *testing.T) {
	ctx, mock := testContext()

	want := errors.New("x")

	calledGet := mock.servers.Repos.MockGet(t, 1)
	var calledVCSRepoResolveRevision bool
	mock.stores.RepoVCS.MockOpen(t, 1, vcstest.MockRepository{
		ResolveRevision_: func(rev string) (vcs.CommitID, error) {
			calledVCSRepoResolveRevision = true
			return "", errors.New("x")
		},
	})

	_, err := resolveRepoRev(ctx, 1, strings.Repeat("a", 40))
	if !reflect.DeepEqual(err, want) {
		t.Fatalf("got err %v, want %v", err, want)
	}
	if !*calledGet {
		t.Error("!calledGet")
	}
	if !calledVCSRepoResolveRevision {
		t.Error("!calledVCSRepoResolveRevision")
	}
}

func Test_Repos_ListCommits(t *testing.T) {
	wantCommits := []*vcs.Commit{
		{ID: vcs.CommitID("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
		{ID: vcs.CommitID("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")},
		{ID: vcs.CommitID("cccccccccccccccccccccccccccccccccccccccc")},
		{ID: vcs.CommitID("dddddddddddddddddddddddddddddddddddddddd")},
	}

	var s repos
	ctx, mock := testContext()

	calledGet := mock.servers.Repos.MockGet(t, 1)
	mockRepo := vcstest.MockRepository{}
	mockRepo.ResolveRevision_ = func(spec string) (vcs.CommitID, error) {
		if spec != "v" {
			t.Fatalf("call to ResolveRevision with unexpected argument spec=%s", spec)
		}
		return wantCommits[0].ID, nil
	}
	mockRepo.Commits_ = func(opt vcs.CommitsOptions) ([]*vcs.Commit, uint, error) {
		if !(opt.Head == wantCommits[0].ID && opt.Base == "") {
			t.Fatalf("call to Commits with unexpected argument opt=%+v", opt)
		}
		return wantCommits, uint(len(wantCommits)), nil
	}
	mock.stores.RepoVCS.Open_ = func(ctx context.Context, repo int32) (vcs.Repository, error) {
		return mockRepo, nil
	}

	commitList, err := s.ListCommits(ctx, &sourcegraph.ReposListCommitsOp{
		Repo: 1,
		Opt:  &sourcegraph.RepoListCommitsOptions{Head: "v"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantCommits, commitList.Commits) {
		t.Errorf("want %+v, got %+v", wantCommits, commitList.Commits)
	}
	if !*calledGet {
		t.Error("!calledGet")
	}
}
