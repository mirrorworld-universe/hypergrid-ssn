package keeper

import (
	"context"

	"hypergrid-ssn/x/hypergridssn/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FeeSettlementBillAll(ctx context.Context, req *types.QueryAllFeeSettlementBillRequest) (*types.QueryAllFeeSettlementBillResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var feeSettlementBills []types.FeeSettlementBill

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	feeSettlementBillStore := prefix.NewStore(store, types.KeyPrefix(types.FeeSettlementBillKey))

	pageRes, err := query.Paginate(feeSettlementBillStore, req.Pagination, func(key []byte, value []byte) error {
		var feeSettlementBill types.FeeSettlementBill
		if err := k.cdc.Unmarshal(value, &feeSettlementBill); err != nil {
			return err
		}

		feeSettlementBills = append(feeSettlementBills, feeSettlementBill)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFeeSettlementBillResponse{FeeSettlementBill: feeSettlementBills, Pagination: pageRes}, nil
}

func (k Keeper) FeeSettlementBill(ctx context.Context, req *types.QueryGetFeeSettlementBillRequest) (*types.QueryGetFeeSettlementBillResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	feeSettlementBill, found := k.GetFeeSettlementBill(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetFeeSettlementBillResponse{FeeSettlementBill: feeSettlementBill}, nil
}
