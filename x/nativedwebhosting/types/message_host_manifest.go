package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgHostManifest = "host_manifest"

var _ sdk.Msg = &MsgHostManifest{}

func NewMsgHostManifest(creator string, version string, hashFunction string) *MsgHostManifest {
	return &MsgHostManifest{
		Creator:      creator,
		Version:      version,
		HashFunction: hashFunction,
	}
}

func (msg *MsgHostManifest) Route() string {
	return RouterKey
}

func (msg *MsgHostManifest) Type() string {
	return TypeMsgHostManifest
}

func (msg *MsgHostManifest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgHostManifest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgHostManifest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
