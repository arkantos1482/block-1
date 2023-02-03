package blocktest1_test

import (
	"testing"

	keepertest "github.com/arkantos1482/block-test-1/testutil/keeper"
	"github.com/arkantos1482/block-test-1/testutil/nullify"
	"github.com/arkantos1482/block-test-1/x/blocktest1"
	"github.com/arkantos1482/block-test-1/x/blocktest1/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Blocktest1Keeper(t)
	blocktest1.InitGenesis(ctx, *k, genesisState)
	got := blocktest1.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
