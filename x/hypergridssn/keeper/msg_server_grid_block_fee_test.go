package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"hypergrid-ssn/x/hypergridssn/types"
)

func TestGridBlockFeeMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateGridBlockFee(wctx, &types.MsgCreateGridBlockFee{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}
