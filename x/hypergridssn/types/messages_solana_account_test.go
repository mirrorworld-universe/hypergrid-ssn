package types

import (
	"testing"

	"hypergrid-ssn/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateSolanaAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateSolanaAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateSolanaAccount{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateSolanaAccount{
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

func TestMsgUpdateSolanaAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateSolanaAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateSolanaAccount{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateSolanaAccount{
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

func TestMsgDeleteSolanaAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteSolanaAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteSolanaAccount{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteSolanaAccount{
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
