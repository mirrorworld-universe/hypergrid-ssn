package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "hypergrid-ssn/testutil/keeper"
	"hypergrid-ssn/x/hypergridssn/keeper"
	"hypergrid-ssn/x/hypergridssn/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestHypergridNodeMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.HypergridssnKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateHypergridNode{Creator: creator,
			Pubkey: strconv.Itoa(i),
		}
		_, err := srv.CreateHypergridNode(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetHypergridNode(ctx,
			expected.Pubkey,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestHypergridNodeMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateHypergridNode
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateHypergridNode{Creator: creator,
				Pubkey: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateHypergridNode{Creator: "B",
				Pubkey: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateHypergridNode{Creator: creator,
				Pubkey: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.HypergridssnKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateHypergridNode{Creator: creator,
				Pubkey: strconv.Itoa(0),
			}
			_, err := srv.CreateHypergridNode(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateHypergridNode(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetHypergridNode(ctx,
					expected.Pubkey,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestHypergridNodeMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteHypergridNode
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteHypergridNode{Creator: creator,
				Pubkey: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteHypergridNode{Creator: "B",
				Pubkey: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteHypergridNode{Creator: creator,
				Pubkey: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.HypergridssnKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateHypergridNode(ctx, &types.MsgCreateHypergridNode{Creator: creator,
				Pubkey: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteHypergridNode(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetHypergridNode(ctx,
					tc.request.Pubkey,
				)
				require.False(t, found)
			}
		})
	}
}
