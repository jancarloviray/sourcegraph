// GENERATED CODE - DO NOT EDIT!
//
// Generated by:
//
//   go run gen_remote.go
//
// Called via:
//
//   go generate
//

package remote

import (
	"golang.org/x/net/context"
	"sourcegraph.com/sourcegraph/go-vcs/vcs"
	"sourcegraph.com/sourcegraph/srclib/unit"
	"sourcegraph.com/sqs/pbtypes"
	"src.sourcegraph.com/sourcegraph/go-sourcegraph/sourcegraph"
	"src.sourcegraph.com/sourcegraph/pkg/inventory"
	"src.sourcegraph.com/sourcegraph/svc"
)

// Services is a full set of remote services (implemented by calling a client to invoke each method on a remote server).
var Services = svc.Services{
	Accounts:            remoteAccounts{},
	Auth:                remoteAuth{},
	Builds:              remoteBuilds{},
	Changesets:          remoteChangesets{},
	Defs:                remoteDefs{},
	Deltas:              remoteDeltas{},
	GraphUplink:         remoteGraphUplink{},
	Markdown:            remoteMarkdown{},
	Meta:                remoteMeta{},
	MirrorRepos:         remoteMirrorRepos{},
	MirroredRepoSSHKeys: remoteMirroredRepoSSHKeys{},
	Notify:              remoteNotify{},
	Orgs:                remoteOrgs{},
	People:              remotePeople{},
	RegisteredClients:   remoteRegisteredClients{},
	RepoBadges:          remoteRepoBadges{},
	RepoStatuses:        remoteRepoStatuses{},
	RepoTree:            remoteRepoTree{},
	Repos:               remoteRepos{},
	Search:              remoteSearch{},
	Storage:             remoteStorage{},
	Units:               remoteUnits{},
	UserKeys:            remoteUserKeys{},
	Users:               remoteUsers{},
}

type remoteAccounts struct{ sourcegraph.AccountsServer }

func (s remoteAccounts) Create(ctx context.Context, v1 *sourcegraph.NewAccount) (*sourcegraph.UserSpec, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Accounts.Create(ctx, v1)
}

func (s remoteAccounts) RequestPasswordReset(ctx context.Context, v1 *sourcegraph.PersonSpec) (*sourcegraph.PendingPasswordReset, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Accounts.RequestPasswordReset(ctx, v1)
}

func (s remoteAccounts) ResetPassword(ctx context.Context, v1 *sourcegraph.NewPassword) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Accounts.ResetPassword(ctx, v1)
}

func (s remoteAccounts) Update(ctx context.Context, v1 *sourcegraph.User) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Accounts.Update(ctx, v1)
}

func (s remoteAccounts) Invite(ctx context.Context, v1 *sourcegraph.AccountInvite) (*sourcegraph.PendingInvite, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Accounts.Invite(ctx, v1)
}

func (s remoteAccounts) AcceptInvite(ctx context.Context, v1 *sourcegraph.AcceptedInvite) (*sourcegraph.UserSpec, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Accounts.AcceptInvite(ctx, v1)
}

func (s remoteAccounts) ListInvites(ctx context.Context, v1 *pbtypes.Void) (*sourcegraph.AccountInviteList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Accounts.ListInvites(ctx, v1)
}

func (s remoteAccounts) DeleteInvite(ctx context.Context, v1 *sourcegraph.InviteSpec) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Accounts.DeleteInvite(ctx, v1)
}

func (s remoteAccounts) Delete(ctx context.Context, v1 *sourcegraph.PersonSpec) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Accounts.Delete(ctx, v1)
}

type remoteAuth struct{ sourcegraph.AuthServer }

func (s remoteAuth) GetAuthorizationCode(ctx context.Context, v1 *sourcegraph.AuthorizationCodeRequest) (*sourcegraph.AuthorizationCode, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Auth.GetAuthorizationCode(ctx, v1)
}

func (s remoteAuth) GetAccessToken(ctx context.Context, v1 *sourcegraph.AccessTokenRequest) (*sourcegraph.AccessTokenResponse, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Auth.GetAccessToken(ctx, v1)
}

func (s remoteAuth) Identify(ctx context.Context, v1 *pbtypes.Void) (*sourcegraph.AuthInfo, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Auth.Identify(ctx, v1)
}

type remoteBuilds struct{ sourcegraph.BuildsServer }

func (s remoteBuilds) Get(ctx context.Context, v1 *sourcegraph.BuildSpec) (*sourcegraph.Build, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Builds.Get(ctx, v1)
}

func (s remoteBuilds) GetRepoBuild(ctx context.Context, v1 *sourcegraph.RepoRevSpec) (*sourcegraph.Build, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Builds.GetRepoBuild(ctx, v1)
}

func (s remoteBuilds) List(ctx context.Context, v1 *sourcegraph.BuildListOptions) (*sourcegraph.BuildList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Builds.List(ctx, v1)
}

func (s remoteBuilds) Create(ctx context.Context, v1 *sourcegraph.BuildsCreateOp) (*sourcegraph.Build, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Builds.Create(ctx, v1)
}

func (s remoteBuilds) Update(ctx context.Context, v1 *sourcegraph.BuildsUpdateOp) (*sourcegraph.Build, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Builds.Update(ctx, v1)
}

func (s remoteBuilds) ListBuildTasks(ctx context.Context, v1 *sourcegraph.BuildsListBuildTasksOp) (*sourcegraph.BuildTaskList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Builds.ListBuildTasks(ctx, v1)
}

func (s remoteBuilds) CreateTasks(ctx context.Context, v1 *sourcegraph.BuildsCreateTasksOp) (*sourcegraph.BuildTaskList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Builds.CreateTasks(ctx, v1)
}

func (s remoteBuilds) UpdateTask(ctx context.Context, v1 *sourcegraph.BuildsUpdateTaskOp) (*sourcegraph.BuildTask, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Builds.UpdateTask(ctx, v1)
}

func (s remoteBuilds) GetTaskLog(ctx context.Context, v1 *sourcegraph.BuildsGetTaskLogOp) (*sourcegraph.LogEntries, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Builds.GetTaskLog(ctx, v1)
}

func (s remoteBuilds) DequeueNext(ctx context.Context, v1 *sourcegraph.BuildsDequeueNextOp) (*sourcegraph.Build, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Builds.DequeueNext(ctx, v1)
}

type remoteChangesets struct{ sourcegraph.ChangesetsServer }

func (s remoteChangesets) Create(ctx context.Context, v1 *sourcegraph.ChangesetCreateOp) (*sourcegraph.Changeset, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Changesets.Create(ctx, v1)
}

func (s remoteChangesets) Get(ctx context.Context, v1 *sourcegraph.ChangesetSpec) (*sourcegraph.Changeset, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Changesets.Get(ctx, v1)
}

func (s remoteChangesets) List(ctx context.Context, v1 *sourcegraph.ChangesetListOp) (*sourcegraph.ChangesetList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Changesets.List(ctx, v1)
}

func (s remoteChangesets) Update(ctx context.Context, v1 *sourcegraph.ChangesetUpdateOp) (*sourcegraph.ChangesetEvent, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Changesets.Update(ctx, v1)
}

func (s remoteChangesets) Merge(ctx context.Context, v1 *sourcegraph.ChangesetMergeOp) (*sourcegraph.ChangesetEvent, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Changesets.Merge(ctx, v1)
}

func (s remoteChangesets) UpdateAffected(ctx context.Context, v1 *sourcegraph.ChangesetUpdateAffectedOp) (*sourcegraph.ChangesetEventList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Changesets.UpdateAffected(ctx, v1)
}

func (s remoteChangesets) CreateReview(ctx context.Context, v1 *sourcegraph.ChangesetCreateReviewOp) (*sourcegraph.ChangesetReview, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Changesets.CreateReview(ctx, v1)
}

func (s remoteChangesets) ListReviews(ctx context.Context, v1 *sourcegraph.ChangesetListReviewsOp) (*sourcegraph.ChangesetReviewList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Changesets.ListReviews(ctx, v1)
}

func (s remoteChangesets) ListEvents(ctx context.Context, v1 *sourcegraph.ChangesetSpec) (*sourcegraph.ChangesetEventList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Changesets.ListEvents(ctx, v1)
}

type remoteDefs struct{ sourcegraph.DefsServer }

func (s remoteDefs) Get(ctx context.Context, v1 *sourcegraph.DefsGetOp) (*sourcegraph.Def, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Defs.Get(ctx, v1)
}

func (s remoteDefs) List(ctx context.Context, v1 *sourcegraph.DefListOptions) (*sourcegraph.DefList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Defs.List(ctx, v1)
}

func (s remoteDefs) ListRefs(ctx context.Context, v1 *sourcegraph.DefsListRefsOp) (*sourcegraph.RefList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Defs.ListRefs(ctx, v1)
}

func (s remoteDefs) ListExamples(ctx context.Context, v1 *sourcegraph.DefsListExamplesOp) (*sourcegraph.ExampleList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Defs.ListExamples(ctx, v1)
}

func (s remoteDefs) ListAuthors(ctx context.Context, v1 *sourcegraph.DefsListAuthorsOp) (*sourcegraph.DefAuthorList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Defs.ListAuthors(ctx, v1)
}

func (s remoteDefs) ListClients(ctx context.Context, v1 *sourcegraph.DefsListClientsOp) (*sourcegraph.DefClientList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Defs.ListClients(ctx, v1)
}

type remoteDeltas struct{ sourcegraph.DeltasServer }

func (s remoteDeltas) Get(ctx context.Context, v1 *sourcegraph.DeltaSpec) (*sourcegraph.Delta, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Deltas.Get(ctx, v1)
}

func (s remoteDeltas) ListUnits(ctx context.Context, v1 *sourcegraph.DeltasListUnitsOp) (*sourcegraph.UnitDeltaList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Deltas.ListUnits(ctx, v1)
}

func (s remoteDeltas) ListDefs(ctx context.Context, v1 *sourcegraph.DeltasListDefsOp) (*sourcegraph.DeltaDefs, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Deltas.ListDefs(ctx, v1)
}

func (s remoteDeltas) ListFiles(ctx context.Context, v1 *sourcegraph.DeltasListFilesOp) (*sourcegraph.DeltaFiles, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Deltas.ListFiles(ctx, v1)
}

func (s remoteDeltas) ListAffectedAuthors(ctx context.Context, v1 *sourcegraph.DeltasListAffectedAuthorsOp) (*sourcegraph.DeltaAffectedPersonList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Deltas.ListAffectedAuthors(ctx, v1)
}

func (s remoteDeltas) ListAffectedClients(ctx context.Context, v1 *sourcegraph.DeltasListAffectedClientsOp) (*sourcegraph.DeltaAffectedPersonList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Deltas.ListAffectedClients(ctx, v1)
}

type remoteGraphUplink struct{ sourcegraph.GraphUplinkServer }

func (s remoteGraphUplink) Push(ctx context.Context, v1 *sourcegraph.MetricsSnapshot) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.GraphUplink.Push(ctx, v1)
}

func (s remoteGraphUplink) PushEvents(ctx context.Context, v1 *sourcegraph.UserEventList) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.GraphUplink.PushEvents(ctx, v1)
}

type remoteMarkdown struct{ sourcegraph.MarkdownServer }

func (s remoteMarkdown) Render(ctx context.Context, v1 *sourcegraph.MarkdownRenderOp) (*sourcegraph.MarkdownData, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Markdown.Render(ctx, v1)
}

type remoteMeta struct{ sourcegraph.MetaServer }

func (s remoteMeta) Status(ctx context.Context, v1 *pbtypes.Void) (*sourcegraph.ServerStatus, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Meta.Status(ctx, v1)
}

func (s remoteMeta) Config(ctx context.Context, v1 *pbtypes.Void) (*sourcegraph.ServerConfig, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Meta.Config(ctx, v1)
}

type remoteMirrorRepos struct{ sourcegraph.MirrorReposServer }

func (s remoteMirrorRepos) RefreshVCS(ctx context.Context, v1 *sourcegraph.MirrorReposRefreshVCSOp) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.MirrorRepos.RefreshVCS(ctx, v1)
}

type remoteMirroredRepoSSHKeys struct {
	sourcegraph.MirroredRepoSSHKeysServer
}

func (s remoteMirroredRepoSSHKeys) Create(ctx context.Context, v1 *sourcegraph.MirroredRepoSSHKeysCreateOp) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.MirroredRepoSSHKeys.Create(ctx, v1)
}

func (s remoteMirroredRepoSSHKeys) Get(ctx context.Context, v1 *sourcegraph.RepoSpec) (*sourcegraph.SSHPrivateKey, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.MirroredRepoSSHKeys.Get(ctx, v1)
}

func (s remoteMirroredRepoSSHKeys) Delete(ctx context.Context, v1 *sourcegraph.RepoSpec) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.MirroredRepoSSHKeys.Delete(ctx, v1)
}

type remoteNotify struct{ sourcegraph.NotifyServer }

func (s remoteNotify) GenericEvent(ctx context.Context, v1 *sourcegraph.NotifyGenericEvent) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Notify.GenericEvent(ctx, v1)
}

type remoteOrgs struct{ sourcegraph.OrgsServer }

func (s remoteOrgs) Get(ctx context.Context, v1 *sourcegraph.OrgSpec) (*sourcegraph.Org, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Orgs.Get(ctx, v1)
}

func (s remoteOrgs) List(ctx context.Context, v1 *sourcegraph.OrgsListOp) (*sourcegraph.OrgList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Orgs.List(ctx, v1)
}

func (s remoteOrgs) ListMembers(ctx context.Context, v1 *sourcegraph.OrgsListMembersOp) (*sourcegraph.UserList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Orgs.ListMembers(ctx, v1)
}

type remotePeople struct{ sourcegraph.PeopleServer }

func (s remotePeople) Get(ctx context.Context, v1 *sourcegraph.PersonSpec) (*sourcegraph.Person, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.People.Get(ctx, v1)
}

type remoteRegisteredClients struct {
	sourcegraph.RegisteredClientsServer
}

func (s remoteRegisteredClients) Get(ctx context.Context, v1 *sourcegraph.RegisteredClientSpec) (*sourcegraph.RegisteredClient, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RegisteredClients.Get(ctx, v1)
}

func (s remoteRegisteredClients) GetCurrent(ctx context.Context, v1 *pbtypes.Void) (*sourcegraph.RegisteredClient, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RegisteredClients.GetCurrent(ctx, v1)
}

func (s remoteRegisteredClients) Create(ctx context.Context, v1 *sourcegraph.RegisteredClient) (*sourcegraph.RegisteredClient, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RegisteredClients.Create(ctx, v1)
}

func (s remoteRegisteredClients) Update(ctx context.Context, v1 *sourcegraph.RegisteredClient) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RegisteredClients.Update(ctx, v1)
}

func (s remoteRegisteredClients) Delete(ctx context.Context, v1 *sourcegraph.RegisteredClientSpec) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RegisteredClients.Delete(ctx, v1)
}

func (s remoteRegisteredClients) List(ctx context.Context, v1 *sourcegraph.RegisteredClientListOptions) (*sourcegraph.RegisteredClientList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RegisteredClients.List(ctx, v1)
}

func (s remoteRegisteredClients) GetUserPermissions(ctx context.Context, v1 *sourcegraph.UserPermissionsOptions) (*sourcegraph.UserPermissions, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RegisteredClients.GetUserPermissions(ctx, v1)
}

func (s remoteRegisteredClients) SetUserPermissions(ctx context.Context, v1 *sourcegraph.UserPermissions) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RegisteredClients.SetUserPermissions(ctx, v1)
}

func (s remoteRegisteredClients) ListUserPermissions(ctx context.Context, v1 *sourcegraph.RegisteredClientSpec) (*sourcegraph.UserPermissionsList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RegisteredClients.ListUserPermissions(ctx, v1)
}

type remoteRepoBadges struct{ sourcegraph.RepoBadgesServer }

func (s remoteRepoBadges) ListBadges(ctx context.Context, v1 *sourcegraph.RepoSpec) (*sourcegraph.BadgeList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RepoBadges.ListBadges(ctx, v1)
}

func (s remoteRepoBadges) ListCounters(ctx context.Context, v1 *sourcegraph.RepoSpec) (*sourcegraph.CounterList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RepoBadges.ListCounters(ctx, v1)
}

func (s remoteRepoBadges) RecordHit(ctx context.Context, v1 *sourcegraph.RepoSpec) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RepoBadges.RecordHit(ctx, v1)
}

func (s remoteRepoBadges) CountHits(ctx context.Context, v1 *sourcegraph.RepoBadgesCountHitsOp) (*sourcegraph.RepoBadgesCountHitsResult, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RepoBadges.CountHits(ctx, v1)
}

type remoteRepoStatuses struct{ sourcegraph.RepoStatusesServer }

func (s remoteRepoStatuses) GetCombined(ctx context.Context, v1 *sourcegraph.RepoRevSpec) (*sourcegraph.CombinedStatus, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RepoStatuses.GetCombined(ctx, v1)
}

func (s remoteRepoStatuses) Create(ctx context.Context, v1 *sourcegraph.RepoStatusesCreateOp) (*sourcegraph.RepoStatus, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RepoStatuses.Create(ctx, v1)
}

type remoteRepoTree struct{ sourcegraph.RepoTreeServer }

func (s remoteRepoTree) Get(ctx context.Context, v1 *sourcegraph.RepoTreeGetOp) (*sourcegraph.TreeEntry, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RepoTree.Get(ctx, v1)
}

func (s remoteRepoTree) Search(ctx context.Context, v1 *sourcegraph.RepoTreeSearchOp) (*sourcegraph.VCSSearchResultList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RepoTree.Search(ctx, v1)
}

func (s remoteRepoTree) List(ctx context.Context, v1 *sourcegraph.RepoTreeListOp) (*sourcegraph.RepoTreeListResult, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.RepoTree.List(ctx, v1)
}

type remoteRepos struct{ sourcegraph.ReposServer }

func (s remoteRepos) Get(ctx context.Context, v1 *sourcegraph.RepoSpec) (*sourcegraph.Repo, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.Get(ctx, v1)
}

func (s remoteRepos) List(ctx context.Context, v1 *sourcegraph.RepoListOptions) (*sourcegraph.RepoList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.List(ctx, v1)
}

func (s remoteRepos) Create(ctx context.Context, v1 *sourcegraph.ReposCreateOp) (*sourcegraph.Repo, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.Create(ctx, v1)
}

func (s remoteRepos) Update(ctx context.Context, v1 *sourcegraph.ReposUpdateOp) (*sourcegraph.Repo, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.Update(ctx, v1)
}

func (s remoteRepos) Delete(ctx context.Context, v1 *sourcegraph.RepoSpec) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.Delete(ctx, v1)
}

func (s remoteRepos) GetReadme(ctx context.Context, v1 *sourcegraph.RepoRevSpec) (*sourcegraph.Readme, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.GetReadme(ctx, v1)
}

func (s remoteRepos) GetConfig(ctx context.Context, v1 *sourcegraph.RepoSpec) (*sourcegraph.RepoConfig, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.GetConfig(ctx, v1)
}

func (s remoteRepos) GetCommit(ctx context.Context, v1 *sourcegraph.RepoRevSpec) (*vcs.Commit, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.GetCommit(ctx, v1)
}

func (s remoteRepos) ListCommits(ctx context.Context, v1 *sourcegraph.ReposListCommitsOp) (*sourcegraph.CommitList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.ListCommits(ctx, v1)
}

func (s remoteRepos) ListBranches(ctx context.Context, v1 *sourcegraph.ReposListBranchesOp) (*sourcegraph.BranchList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.ListBranches(ctx, v1)
}

func (s remoteRepos) ListTags(ctx context.Context, v1 *sourcegraph.ReposListTagsOp) (*sourcegraph.TagList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.ListTags(ctx, v1)
}

func (s remoteRepos) ListCommitters(ctx context.Context, v1 *sourcegraph.ReposListCommittersOp) (*sourcegraph.CommitterList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.ListCommitters(ctx, v1)
}

func (s remoteRepos) GetSrclibDataVersionForPath(ctx context.Context, v1 *sourcegraph.TreeEntrySpec) (*sourcegraph.SrclibDataVersion, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.GetSrclibDataVersionForPath(ctx, v1)
}

func (s remoteRepos) ConfigureApp(ctx context.Context, v1 *sourcegraph.RepoConfigureAppOp) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.ConfigureApp(ctx, v1)
}

func (s remoteRepos) GetInventory(ctx context.Context, v1 *sourcegraph.RepoRevSpec) (*inventory.Inventory, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Repos.GetInventory(ctx, v1)
}

type remoteSearch struct{ sourcegraph.SearchServer }

func (s remoteSearch) SearchTokens(ctx context.Context, v1 *sourcegraph.TokenSearchOptions) (*sourcegraph.DefList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Search.SearchTokens(ctx, v1)
}

func (s remoteSearch) SearchText(ctx context.Context, v1 *sourcegraph.TextSearchOptions) (*sourcegraph.VCSSearchResultList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Search.SearchText(ctx, v1)
}

type remoteStorage struct{ sourcegraph.StorageServer }

func (s remoteStorage) Get(ctx context.Context, v1 *sourcegraph.StorageKey) (*sourcegraph.StorageValue, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Storage.Get(ctx, v1)
}

func (s remoteStorage) Put(ctx context.Context, v1 *sourcegraph.StoragePutOp) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Storage.Put(ctx, v1)
}

func (s remoteStorage) PutNoOverwrite(ctx context.Context, v1 *sourcegraph.StoragePutOp) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Storage.PutNoOverwrite(ctx, v1)
}

func (s remoteStorage) Delete(ctx context.Context, v1 *sourcegraph.StorageKey) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Storage.Delete(ctx, v1)
}

func (s remoteStorage) Exists(ctx context.Context, v1 *sourcegraph.StorageKey) (*sourcegraph.StorageExists, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Storage.Exists(ctx, v1)
}

func (s remoteStorage) List(ctx context.Context, v1 *sourcegraph.StorageKey) (*sourcegraph.StorageList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Storage.List(ctx, v1)
}

type remoteUnits struct{ sourcegraph.UnitsServer }

func (s remoteUnits) Get(ctx context.Context, v1 *sourcegraph.UnitSpec) (*unit.RepoSourceUnit, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Units.Get(ctx, v1)
}

func (s remoteUnits) List(ctx context.Context, v1 *sourcegraph.UnitListOptions) (*sourcegraph.RepoSourceUnitList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Units.List(ctx, v1)
}

type remoteUserKeys struct{ sourcegraph.UserKeysServer }

func (s remoteUserKeys) AddKey(ctx context.Context, v1 *sourcegraph.SSHPublicKey) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.UserKeys.AddKey(ctx, v1)
}

func (s remoteUserKeys) LookupUser(ctx context.Context, v1 *sourcegraph.SSHPublicKey) (*sourcegraph.UserSpec, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.UserKeys.LookupUser(ctx, v1)
}

func (s remoteUserKeys) DeleteKey(ctx context.Context, v1 *sourcegraph.SSHPublicKey) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.UserKeys.DeleteKey(ctx, v1)
}

func (s remoteUserKeys) ListKeys(ctx context.Context, v1 *pbtypes.Void) (*sourcegraph.SSHKeyList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.UserKeys.ListKeys(ctx, v1)
}

func (s remoteUserKeys) DeleteAllKeys(ctx context.Context, v1 *pbtypes.Void) (*pbtypes.Void, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.UserKeys.DeleteAllKeys(ctx, v1)
}

type remoteUsers struct{ sourcegraph.UsersServer }

func (s remoteUsers) Get(ctx context.Context, v1 *sourcegraph.UserSpec) (*sourcegraph.User, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Users.Get(ctx, v1)
}

func (s remoteUsers) GetWithEmail(ctx context.Context, v1 *sourcegraph.EmailAddr) (*sourcegraph.User, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Users.GetWithEmail(ctx, v1)
}

func (s remoteUsers) ListEmails(ctx context.Context, v1 *sourcegraph.UserSpec) (*sourcegraph.EmailAddrList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Users.ListEmails(ctx, v1)
}

func (s remoteUsers) List(ctx context.Context, v1 *sourcegraph.UsersListOptions) (*sourcegraph.UserList, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Users.List(ctx, v1)
}

func (s remoteUsers) Count(ctx context.Context, v1 *pbtypes.Void) (*sourcegraph.UserCount, error) {
	c, err := sourcegraph.NewClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return c.Users.Count(ctx, v1)
}
