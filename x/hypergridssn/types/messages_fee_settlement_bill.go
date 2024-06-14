package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateFeeSettlementBill{}

func NewMsgCreateFeeSettlementBill(creator string, fromId uint64, endId uint64, bill string, status int32) *MsgCreateFeeSettlementBill {
	return &MsgCreateFeeSettlementBill{
		Creator: creator,
		FromId:  fromId,
		EndId:   endId,
	}
}

func (msg *MsgCreateFeeSettlementBill) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
