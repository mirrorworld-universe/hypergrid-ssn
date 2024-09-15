package keeper

import (
	"context"

	"hypergrid-ssn/x/hypergridssn/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateGridBlockFee(goCtx context.Context, msg *types.MsgCreateGridBlockFee) (*types.MsgCreateGridBlockFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var n = len(msg.Items)
	if n == 0 {
		return nil, errorsmod.Wrap(types.ErrInvalidLengthGridBlockFee, "no GridBlockFeeItem")
	} else if n > 10000 {
		return nil, errorsmod.Wrap(types.ErrIntOverflowGridBlockFee, "too many GridBlockFeeItem")
	}

	ids := make([]uint64, 0, len(msg.Items))
	for _, item := range msg.Items {
		var gridBlockFee = types.GridBlockFee{
			Creator:   msg.Creator,
			Grid:      item.Grid,
			Slot:      item.Slot,
			Blockhash: item.Blockhash,
			Blocktime: item.Blocktime,
			Fee:       item.Fee,
		}

		//make blockhash unique
		if k.HasGridBlockFeeHash(ctx, gridBlockFee.Blockhash) {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "blockhash already")
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
