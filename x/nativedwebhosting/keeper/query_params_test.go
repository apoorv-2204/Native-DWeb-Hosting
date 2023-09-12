package keeper_test

import (
	"testing"

	testkeeper "github.com/apoorv-2204/Native-DWeb-Hosting/testutil/keeper"
	"github.com/apoorv-2204/Native-DWeb-Hosting/x/nativedwebhosting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.NativedwebhostingKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
