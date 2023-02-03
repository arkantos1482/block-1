package keeper_test

import (
	"testing"

	testkeeper "github.com/arkantos1482/block-test-1/testutil/keeper"
	"github.com/arkantos1482/block-test-1/x/blocktest1/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Blocktest1Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
