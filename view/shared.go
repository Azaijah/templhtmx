package view

import (
	"context"

	"github.com/templhtmx/types"
)

func AuthenticatedUser(ctx context.Context) types.AuthenticatedUser {

	user, ok := ctx.Value(types.UserCtxKey).(types.AuthenticatedUser)

	if !ok {
		return types.AuthenticatedUser{}
	}

	return user
}
