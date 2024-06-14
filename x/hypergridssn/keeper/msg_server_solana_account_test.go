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

func TestSolanaAccountMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.HypergridssnKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateSolanaAccount{Creator: creator,
			Address: strconv.Itoa(i),
			Version: strconv.Itoa(i),
			Source:  strconv.Itoa(i),
		}
		_, err := srv.CreateSolanaAccount(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetSolanaAccount(ctx,
			expected.Address,
			expected.Version,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestSolanaAccountMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateSolanaAccount
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateSolanaAccount{Creator: creator,
				Address: strconv.Itoa(0),
				Version: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateSolanaAccount{Creator: "B",
				Address: strconv.Itoa(0),
				Version: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateSolanaAccount{Creator: creator,
				Address: strconv.Itoa(100000),
				Version: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.HypergridssnKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateSolanaAccount{Creator: creator,
				Address: strconv.Itoa(0),
				Version: strconv.Itoa(0),
				Source:  strconv.Itoa(0),
			}
			_, err := srv.CreateSolanaAccount(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateSolanaAccount(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetSolanaAccount(ctx,
					expected.Address,
					expected.Version,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestSolanaAccountMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteSolanaAccount
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteSolanaAccount{Creator: creator,
				Address: strconv.Itoa(0),
				Version: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteSolanaAccount{Creator: "B",
				Address: strconv.Itoa(0),
				Version: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteSolanaAccount{Creator: creator,
				Address: strconv.Itoa(100000),
				Version: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.HypergridssnKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateSolanaAccount(ctx, &types.MsgCreateSolanaAccount{Creator: creator,
				Address: strconv.Itoa(0),
				Version: strconv.Itoa(0),
				Source:  strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteSolanaAccount(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetSolanaAccount(ctx,
					tc.request.Address,
					tc.request.Version,
				)
				require.False(t, found)
			}
		})
	}
}
