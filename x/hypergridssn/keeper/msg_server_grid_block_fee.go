package keeper

import (
	"context"

	"hypergrid-ssn/x/hypergridssn/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGridBlockFee(goCtx context.Context, msg *types.MsgCreateGridBlockFee) (*types.MsgCreateGridBlockFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ids := make([]uint64, len(msg.Items))
	for _, item := range msg.Items {
		var gridBlockFee = types.GridBlockFee{
			Creator:   msg.Creator,
			Grid:      item.Grid,
			Slot:      item.Slot,
			Blockhash: item.Blockhash,
			Blocktime: item.Blocktime,
			Fee:       item.Fee,
		}

		id := k.AppendGridBlockFee(
			ctx,
			gridBlockFee,
		)
		ids = append(ids, id)
	}

	return &types.MsgCreateGridBlockFeeResponse{
		Id: ids[len(ids)-1],
	}, nil
}
