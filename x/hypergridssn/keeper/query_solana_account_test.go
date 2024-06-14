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

func TestSolanaAccountQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	msgs := createNSolanaAccount(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetSolanaAccountRequest
		response *types.QueryGetSolanaAccountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSolanaAccountRequest{
				Address: msgs[0].Address,
				Version: msgs[0].Version,
			},
			response: &types.QueryGetSolanaAccountResponse{SolanaAccount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSolanaAccountRequest{
				Address: msgs[1].Address,
				Version: msgs[1].Version,
			},
			response: &types.QueryGetSolanaAccountResponse{SolanaAccount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSolanaAccountRequest{
				Address: strconv.Itoa(100000),
				Version: strconv.Itoa(100000),
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
			response, err := keeper.SolanaAccount(ctx, tc.request)
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

func TestSolanaAccountQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.HypergridssnKeeper(t)
	msgs := createNSolanaAccount(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSolanaAccountRequest {
		return &types.QueryAllSolanaAccountRequest{
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
			resp, err := keeper.SolanaAccountAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SolanaAccount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SolanaAccount),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SolanaAccountAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SolanaAccount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SolanaAccount),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SolanaAccountAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.SolanaAccount),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SolanaAccountAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
