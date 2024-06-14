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

func (k Keeper) GridBlockFeeAll(ctx context.Context, req *types.QueryAllGridBlockFeeRequest) (*types.QueryAllGridBlockFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var gridBlockFees []types.GridBlockFee

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	gridBlockFeeStore := prefix.NewStore(store, types.KeyPrefix(types.GridBlockFeeKey))

	pageRes, err := query.Paginate(gridBlockFeeStore, req.Pagination, func(key []byte, value []byte) error {
		var gridBlockFee types.GridBlockFee
		if err := k.cdc.Unmarshal(value, &gridBlockFee); err != nil {
			return err
		}

		gridBlockFees = append(gridBlockFees, gridBlockFee)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGridBlockFeeResponse{GridBlockFee: gridBlockFees, Pagination: pageRes}, nil
}

func (k Keeper) GridBlockFee(ctx context.Context, req *types.QueryGetGridBlockFeeRequest) (*types.QueryGetGridBlockFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	gridBlockFee, found := k.GetGridBlockFee(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetGridBlockFeeResponse{GridBlockFee: gridBlockFee}, nil
}
