package keeper

import (
	"context"
	"encoding/json"
	"strconv"

	"hypergrid-ssn/x/hypergridssn/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Bill struct {
	Grid string

	Fee uint64
}

func (k msgServer) CreateFeeSettlementBill(goCtx context.Context, msg *types.MsgCreateFeeSettlementBill) (*types.MsgCreateFeeSettlementBillResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// size := k.GetFeeSettlementBillCount(ctx)
	// val, found := k.GetFeeSettlementBill(ctx, size-1)
	// startId := uint64(0)
	// if found {
	// 	startId = val.EndId
	// }

	startId := msg.FromId

	bills := make(map[string]uint64)
	for i := startId; i < msg.EndId; i++ {
		item, found := k.GetGridBlockFee(ctx, i)
		if !found {
			break
		}

		fee, _ := strconv.ParseUint(item.Fee, 10, 64)
		bills[item.Grid] += fee
	}

	billBytes, err := json.Marshal(bills)
	if err != nil {
		// handle the error, e.g. return an error response or log the error
		return nil, errorsmod.Wrap(err, "failed to marshal bills")
	}

	var feeSettlementBill = types.FeeSettlementBill{
		Creator: msg.Creator,
		FromId:  msg.FromId,
		EndId:   msg.EndId,
		Bill:    string(billBytes),
		Status:  0,
	}

	id := k.AppendFeeSettlementBill(
		ctx,
		feeSettlementBill,
	)

	//todo: call sonic grid to settle the fee

	return &types.MsgCreateFeeSettlementBillResponse{
		Id: id,
	}, nil
}
