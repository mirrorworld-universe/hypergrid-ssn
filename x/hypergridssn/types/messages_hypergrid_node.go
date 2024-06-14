package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHypergridNode{}

func NewMsgCreateHypergridNode(
	creator string,
	pubkey string,
	name string,
	rpc string,
	role int32,
	starttime int32,

) *MsgCreateHypergridNode {
	return &MsgCreateHypergridNode{
		Creator:   creator,
		Pubkey:    pubkey,
		Name:      name,
		Rpc:       rpc,
		Role:      role,
		Starttime: starttime,
	}
}

func (msg *MsgCreateHypergridNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHypergridNode{}

func NewMsgUpdateHypergridNode(
	creator string,
	pubkey string,
	name string,
	rpc string,
	role int32,
	starttime int32,

) *MsgUpdateHypergridNode {
	return &MsgUpdateHypergridNode{
		Creator:   creator,
		Pubkey:    pubkey,
		Name:      name,
		Rpc:       rpc,
		Role:      role,
		Starttime: starttime,
	}
}

func (msg *MsgUpdateHypergridNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteHypergridNode{}

func NewMsgDeleteHypergridNode(
	creator string,
	pubkey string,

) *MsgDeleteHypergridNode {
	return &MsgDeleteHypergridNode{
		Creator: creator,
		Pubkey:  pubkey,
	}
}

func (msg *MsgDeleteHypergridNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
