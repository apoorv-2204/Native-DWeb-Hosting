package keeper_test

import (
	"testing"

	testkeeper "github.com/apoorv-2204/Native-DWeb-Hosting/testutil/keeper"
	"github.com/apoorv-2204/Native-DWeb-Hosting/x/nativedwebhosting/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NativedwebhostingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
