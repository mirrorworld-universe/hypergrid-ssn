package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSolanaAccount{},
		&MsgUpdateSolanaAccount{},
		&MsgDeleteSolanaAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGridBlockFee{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHypergridNode{},
		&MsgUpdateHypergridNode{},
		&MsgDeleteHypergridNode{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateFeeSettlementBill{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
