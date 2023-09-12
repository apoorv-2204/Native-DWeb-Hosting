package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStorage = "storage"

var _ sdk.Msg = &MsgStorage{}

func NewMsgStorage(creator string, version string) *MsgStorage {
	return &MsgStorage{
		Creator: creator,
		Version: version,
	}
}

func (msg *MsgStorage) Route() string {
	return RouterKey
}

func (msg *MsgStorage) Type() string {
	return TypeMsgStorage
}

func (msg *MsgStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
