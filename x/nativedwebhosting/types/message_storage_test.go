package types

import (
	"testing"

	"github.com/apoorv-2204/Native-DWeb-Hosting/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgStorage_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgStorage
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgStorage{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgStorage{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
