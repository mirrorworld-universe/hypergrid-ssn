package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGridBlockFee{}

func NewMsgCreateGridBlockFee(creator string, grid string, slot string, blockhash string, blocktime int32, fee string) *MsgCreateGridBlockFee {
	return &MsgCreateGridBlockFee{
		Creator:   creator,
		Grid:      grid,
		Slot:      slot,
		Blockhash: blockhash,
		Blocktime: blocktime,
		Fee:       fee,
	}
}

func (msg *MsgCreateGridBlockFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
