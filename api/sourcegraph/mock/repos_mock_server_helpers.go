package mock

import (
	"sync"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"sourcegraph.com/sourcegraph/sourcegraph/api/sourcegraph"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/vcs"
)

func (s *ReposServer) MockGet(t *testing.T, wantRepo int32) (called *bool) {
	called = new(bool)
	s.Get_ = func(ctx context.Context, repo *sourcegraph.RepoSpec) (*sourcegraph.Repo, error) {
		*called = true
		if repo.ID != wantRepo {
			t.Errorf("got repo %d, want %d", repo.ID, wantRepo)
			return nil, grpc.Errorf(codes.NotFound, "repo %s not found", wantRepo)
		}
		return &sourcegraph.Repo{ID: repo.ID}, nil
	}
	return
}

func (s *ReposServer) MockGet_Path(t *testing.T, wantRepo int32, repoPath string) (called *bool) {
	called = new(bool)
	s.Get_ = func(ctx context.Context, repo *sourcegraph.RepoSpec) (*sourcegraph.Repo, error) {
		*called = true
		if repo.ID != wantRepo {
			t.Errorf("got repo %d, want %d", repo.ID, wantRepo)
			return nil, grpc.Errorf(codes.NotFound, "repo %s not found", wantRepo)
		}
		return &sourcegraph.Repo{ID: repo.ID, URI: repoPath}, nil
	}
	return
}

func (s *ReposServer) MockGet_Return(t *testing.T, returns *sourcegraph.Repo) (called *bool) {
	called = new(bool)
	s.Get_ = func(ctx context.Context, repo *sourcegraph.RepoSpec) (*sourcegraph.Repo, error) {
		*called = true
		if repo.ID != returns.ID {
			t.Errorf("got repo %d, want %d", repo.ID, returns.ID)
			return nil, grpc.Errorf(codes.NotFound, "repo %d not found", returns.ID)
		}
		return returns, nil
	}
	return
}

func (s *ReposServer) MockResolve_Local(t *testing.T, wantPath string, repoID int32) (called *bool) {
	called = new(bool)
	s.Resolve_ = func(ctx context.Context, op *sourcegraph.RepoResolveOp) (*sourcegraph.RepoResolution, error) {
		*called = true
		if op.Path != wantPath {
			t.Errorf("got repo %q, want %q", op.Path, wantPath)
			return nil, grpc.Errorf(codes.NotFound, "repo path %s resolution failed", wantPath)
		}
		return &sourcegraph.RepoResolution{Repo: repoID, CanonicalPath: wantPath}, nil
	}
	return
}

func (s *ReposServer) MockResolve_Remote(t *testing.T, wantPath string, resolved *sourcegraph.RemoteRepo) (called *bool) {
	called = new(bool)
	s.Resolve_ = func(ctx context.Context, op *sourcegraph.RepoResolveOp) (*sourcegraph.RepoResolution, error) {
		*called = true
		if op.Path != wantPath {
			t.Errorf("got repo %q, want %q", op.Path, wantPath)
			return nil, grpc.Errorf(codes.NotFound, "repo path %s resolution failed", wantPath)
		}
		return &sourcegraph.RepoResolution{RemoteRepo: resolved}, nil
	}
	return
}

func (s *ReposServer) MockResolve_NotFound(t *testing.T, wantPath string) (called *bool) {
	called = new(bool)
	s.Resolve_ = func(ctx context.Context, op *sourcegraph.RepoResolveOp) (*sourcegraph.RepoResolution, error) {
		*called = true
		if op.Path != wantPath {
			t.Errorf("got repo %q, want %q", op.Path, wantPath)
		}
		return nil, grpc.Errorf(codes.NotFound, "repo path %s resolution failed", wantPath)
	}
	return
}

func (s *ReposServer) MockList(t *testing.T, wantRepos ...string) (called *bool) {
	called = new(bool)
	s.List_ = func(ctx context.Context, opt *sourcegraph.RepoListOptions) (*sourcegraph.RepoList, error) {
		*called = true
		repos := make([]*sourcegraph.Repo, len(wantRepos))
		for i, repo := range wantRepos {
			repos[i] = &sourcegraph.Repo{URI: repo}
		}
		return &sourcegraph.RepoList{Repos: repos}, nil
	}
	return
}

func (s *ReposServer) MockListCommits(t *testing.T, wantCommitIDs ...vcs.CommitID) (called *bool) {
	called = new(bool)
	s.ListCommits_ = func(ctx context.Context, op *sourcegraph.ReposListCommitsOp) (*sourcegraph.CommitList, error) {
		*called = true
		commits := make([]*vcs.Commit, len(wantCommitIDs))
		for i, commit := range wantCommitIDs {
			commits[i] = &vcs.Commit{ID: commit}
		}
		return &sourcegraph.CommitList{Commits: commits}, nil
	}
	return
}

func (s *ReposServer) MockResolveRev_NoCheck(t *testing.T, commitID vcs.CommitID) (called *bool) {
	var once sync.Once
	called = new(bool)
	s.ResolveRev_ = func(ctx context.Context, op *sourcegraph.ReposResolveRevOp) (*sourcegraph.ResolvedRev, error) {
		once.Do(func() {
			*called = true
		})
		return &sourcegraph.ResolvedRev{CommitID: string(commitID)}, nil
	}
	return
}

func (s *ReposServer) MockGetCommit_Return_NoCheck(t *testing.T, commit *vcs.Commit) (called *bool) {
	called = new(bool)
	s.GetCommit_ = func(ctx context.Context, repoRev *sourcegraph.RepoRevSpec) (*vcs.Commit, error) {
		*called = true
		return commit, nil
	}
	return
}

func (s *ReposServer) MockGetSrclibDataVersionForPath_Current(t *testing.T) (called *bool) {
	called = new(bool)
	s.GetSrclibDataVersionForPath_ = func(ctx context.Context, entry *sourcegraph.TreeEntrySpec) (*sourcegraph.SrclibDataVersion, error) {
		*called = true
		return &sourcegraph.SrclibDataVersion{CommitID: entry.RepoRev.CommitID}, nil
	}
	return
}
