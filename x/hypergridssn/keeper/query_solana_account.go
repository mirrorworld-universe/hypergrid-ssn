package keeper

import (
	"context"

	"hypergrid-ssn/x/hypergridssn/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SolanaAccountAll(ctx context.Context, req *types.QueryAllSolanaAccountRequest) (*types.QueryAllSolanaAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var solanaAccounts []types.SolanaAccount

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	solanaAccountStore := prefix.NewStore(store, types.KeyPrefix(types.SolanaAccountKeyPrefix))

	pageRes, err := query.Paginate(solanaAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var solanaAccount types.SolanaAccount
		if err := k.cdc.Unmarshal(value, &solanaAccount); err != nil {
			return err
		}

		solanaAccounts = append(solanaAccounts, solanaAccount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSolanaAccountResponse{SolanaAccount: solanaAccounts, Pagination: pageRes}, nil
}

func (k Keeper) SolanaAccount(ctx context.Context, req *types.QueryGetSolanaAccountRequest) (*types.QueryGetSolanaAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetSolanaAccount(
		ctx,
		req.Address,
		req.Version,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSolanaAccountResponse{SolanaAccount: val}, nil
}
