package keeper

import (
	"context"
	// "strconv"

	// solana "hypergrid-ssn/tools"
	"hypergrid-ssn/x/hypergridssn/types"

	errorsmod "cosmossdk.io/errors"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Deprecated: CreateGridInbox
func (k msgServer) CreateGridInbox(goCtx context.Context, msg *types.MsgCreateGridInbox) (*types.MsgCreateGridInboxResponse, error) {
	return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "CreateGridInbox is Deprecated")
	/*
		ctx := sdk.UnwrapSDKContext(goCtx)

		base_layer_rpc := ""
		nodes := k.GetAllHypergridNode(goCtx)
		for _, node := range nodes {
			// get sonic grid node from the list
			if node.Role == 4 {
				base_layer_rpc = node.Rpc
				break
			}
		}

		if base_layer_rpc == "" {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "base layer rpc not found")
		}

		slot, err := strconv.ParseUint(msg.Slot, 10, 64)
		if err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid slot")
		}

		sig, account, err := solana.SendTxInbox(base_layer_rpc, slot, msg.Hash)
		if err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, err.Error())
		}

		println("signature: ", sig)

		var gridInbox = types.GridInbox{
			Creator: msg.Creator,
			Grid:    msg.Grid,
			Account: account.String(),
			Slot:    msg.Slot,
			Hash:    msg.Hash,
		}

		id := k.AppendGridInbox(
			ctx,
			gridInbox,
		)

		return &types.MsgCreateGridInboxResponse{
			Id:     id,
			Txhash: sig.String(),
		}, nil
	*/
}
