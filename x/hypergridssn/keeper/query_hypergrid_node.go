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

func (k Keeper) HypergridNodeAll(ctx context.Context, req *types.QueryAllHypergridNodeRequest) (*types.QueryAllHypergridNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hypergridNodes []types.HypergridNode

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	hypergridNodeStore := prefix.NewStore(store, types.KeyPrefix(types.HypergridNodeKeyPrefix))

	pageRes, err := query.Paginate(hypergridNodeStore, req.Pagination, func(key []byte, value []byte) error {
		var hypergridNode types.HypergridNode
		if err := k.cdc.Unmarshal(value, &hypergridNode); err != nil {
			return err
		}

		hypergridNodes = append(hypergridNodes, hypergridNode)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHypergridNodeResponse{HypergridNode: hypergridNodes, Pagination: pageRes}, nil
}

func (k Keeper) HypergridNode(ctx context.Context, req *types.QueryGetHypergridNodeRequest) (*types.QueryGetHypergridNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetHypergridNode(
		ctx,
		req.Pubkey,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetHypergridNodeResponse{HypergridNode: val}, nil
}
