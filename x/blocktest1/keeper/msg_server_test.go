package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/arkantos1482/block-test-1/testutil/keeper"
	"github.com/arkantos1482/block-test-1/x/blocktest1/keeper"
	"github.com/arkantos1482/block-test-1/x/blocktest1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Blocktest1Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
