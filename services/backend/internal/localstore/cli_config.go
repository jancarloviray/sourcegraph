package localstore

import (
	"golang.org/x/net/context"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/store"
	"sourcegraph.com/sourcegraph/sourcegraph/services/backend/internal/localstore/middleware"
	"sourcegraph.com/sourcegraph/sourcegraph/services/backend/serverctx"
)

func init() {
	stores := store.Stores{
		Accounts:           &accounts{},
		BuildLogs:          &buildLogs{},
		Builds:             &builds{},
		Channel:            &channel{},
		Directory:          &directory{},
		ExternalAuthTokens: &externalAuthTokens{},
		GlobalDefs:         &globalDefs{},
		GlobalDeps:         &globalDeps{},
		GlobalRefs:         &globalRefs{},
		Password:           &password{},
		Queue:              &middleware.InstrumentedQueue{Queue: &queue{}},
		RepoConfigs:        &repoConfigs{},
		RepoStatuses:       &repoStatuses{},
		RepoVCS:            &repoVCS{},
		Repos:              &repos{},
		Users:              &users{},
	}
	serverctx.Funcs = append(serverctx.Funcs, func(ctx context.Context) (context.Context, error) {
		return store.WithStores(ctx, stores), nil
	})
}
