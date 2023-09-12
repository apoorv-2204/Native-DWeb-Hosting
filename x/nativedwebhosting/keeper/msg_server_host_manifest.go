package keeper

import (
	"context"

	"github.com/apoorv-2204/Native-DWeb-Hosting/x/nativedwebhosting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) HostManifest(goCtx context.Context, msg *types.MsgHostManifest) (*types.MsgHostManifestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgHostManifestResponse{}, nil
}
