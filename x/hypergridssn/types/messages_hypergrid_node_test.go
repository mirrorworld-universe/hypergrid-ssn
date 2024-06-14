package types

import (
	"testing"

	"hypergrid-ssn/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateHypergridNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateHypergridNode
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateHypergridNode{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateHypergridNode{
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

func TestMsgUpdateHypergridNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateHypergridNode
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateHypergridNode{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateHypergridNode{
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

func TestMsgDeleteHypergridNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteHypergridNode
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteHypergridNode{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteHypergridNode{
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
