package types

import (
	"testing"

	"hypergrid-ssn/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateGridBlockFee_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateGridBlockFee
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateGridBlockFee{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateGridBlockFee{
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
