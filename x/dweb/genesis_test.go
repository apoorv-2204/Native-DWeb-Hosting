package dweb_test

import (
	"testing"

	keepertest "dweb/testutil/keeper"
	"dweb/testutil/nullify"
	"dweb/x/dweb"
	"dweb/x/dweb/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DwebKeeper(t)
	dweb.InitGenesis(ctx, *k, genesisState)
	got := dweb.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
