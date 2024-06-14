package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "hypergrid-ssn/testutil/keeper"
	"hypergrid-ssn/x/hypergridssn/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.HypergridssnKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
