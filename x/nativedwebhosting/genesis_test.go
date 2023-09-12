package nativedwebhosting_test

import (
	"testing"

	keepertest "github.com/apoorv-2204/Native-DWeb-Hosting/testutil/keeper"
	"github.com/apoorv-2204/Native-DWeb-Hosting/testutil/nullify"
	"github.com/apoorv-2204/Native-DWeb-Hosting/x/nativedwebhosting"
	"github.com/apoorv-2204/Native-DWeb-Hosting/x/nativedwebhosting/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NativedwebhostingKeeper(t)
	nativedwebhosting.InitGenesis(ctx, *k, genesisState)
	got := nativedwebhosting.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
