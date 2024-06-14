package keeper

import (
	"context"

	"hypergrid-ssn/x/hypergridssn/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGridBlockFee(goCtx context.Context, msg *types.MsgCreateGridBlockFee) (*types.MsgCreateGridBlockFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var gridBlockFee = types.GridBlockFee{
		Creator:   msg.Creator,
		Grid:      msg.Grid,
		Slot:      msg.Slot,
		Blockhash: msg.Blockhash,
		Blocktime: msg.Blocktime,
		Fee:       msg.Fee,
	}

	id := k.AppendGridBlockFee(
		ctx,
		gridBlockFee,
	)

	return &types.MsgCreateGridBlockFeeResponse{
		Id: id,
	}, nil
}
