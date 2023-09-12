package keeper

import (
	"context"

	"github.com/apoorv-2204/Native-DWeb-Hosting/x/nativedwebhosting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Storage(goCtx context.Context, msg *types.MsgStorage) (*types.MsgStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgStorageResponse{}, nil
}
