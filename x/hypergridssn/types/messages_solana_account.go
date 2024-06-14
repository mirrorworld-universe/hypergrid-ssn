package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSolanaAccount{}

func NewMsgCreateSolanaAccount(
	creator string,
	address string,
	version string,
	source string,

) *MsgCreateSolanaAccount {
	return &MsgCreateSolanaAccount{
		Creator: creator,
		Address: address,
		Version: version,
		Source:  source,
	}
}

func (msg *MsgCreateSolanaAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSolanaAccount{}

func NewMsgUpdateSolanaAccount(
	creator string,
	address string,
	version string,

) *MsgUpdateSolanaAccount {
	return &MsgUpdateSolanaAccount{
		Creator: creator,
		Address: address,
		Version: version,
	}
}

func (msg *MsgUpdateSolanaAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSolanaAccount{}

func NewMsgDeleteSolanaAccount(
	creator string,
	address string,
	version string,

) *MsgDeleteSolanaAccount {
	return &MsgDeleteSolanaAccount{
		Creator: creator,
		Address: address,
		Version: version,
	}
}

func (msg *MsgDeleteSolanaAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
