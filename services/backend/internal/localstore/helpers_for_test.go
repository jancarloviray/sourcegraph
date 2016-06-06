package localstore

import (
	"golang.org/x/net/context"
	"sourcegraph.com/sourcegraph/sourcegraph/services/backend/accesscontrol"
)

func testContextNoDB() context.Context {
	ctx := context.Background()
	ctx = accesscontrol.WithInsecureSkip(ctx, true)
	return ctx
}
