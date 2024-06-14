package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "hypergrid-ssn/testutil/keeper"
	"hypergrid-ssn/testutil/nullify"
	"hypergrid-ssn/x/hypergridssn/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestHypergridNodeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	msgs := createNHypergridNode(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetHypergridNodeRequest
		response *types.QueryGetHypergridNodeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetHypergridNodeRequest{
				Pubkey: msgs[0].Pubkey,
			},
			response: &types.QueryGetHypergridNodeResponse{HypergridNode: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetHypergridNodeRequest{
				Pubkey: msgs[1].Pubkey,
			},
			response: &types.QueryGetHypergridNodeResponse{HypergridNode: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetHypergridNodeRequest{
				Pubkey: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.HypergridNode(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestHypergridNodeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	msgs := createNHypergridNode(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllHypergridNodeRequest {
		return &types.QueryAllHypergridNodeRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.HypergridNodeAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HypergridNode), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HypergridNode),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.HypergridNodeAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HypergridNode), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HypergridNode),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.HypergridNodeAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.HypergridNode),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.HypergridNodeAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
