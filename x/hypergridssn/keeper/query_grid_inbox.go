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

func (k Keeper) GridInboxAll(ctx context.Context, req *types.QueryAllGridInboxRequest) (*types.QueryAllGridInboxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var gridInboxs []types.GridInbox

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	gridInboxStore := prefix.NewStore(store, types.KeyPrefix(types.GridInboxKey))

	pageRes, err := query.Paginate(gridInboxStore, req.Pagination, func(key []byte, value []byte) error {
		var gridInbox types.GridInbox
		if err := k.cdc.Unmarshal(value, &gridInbox); err != nil {
			return err
		}

		gridInboxs = append(gridInboxs, gridInbox)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGridInboxResponse{GridInbox: gridInboxs, Pagination: pageRes}, nil
}

func (k Keeper) GridInbox(ctx context.Context, req *types.QueryGetGridInboxRequest) (*types.QueryGetGridInboxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	gridInbox, found := k.GetGridInbox(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetGridInboxResponse{GridInbox: gridInbox}, nil
}
