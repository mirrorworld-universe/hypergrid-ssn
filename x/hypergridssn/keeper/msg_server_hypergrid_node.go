package keeper

import (
	"context"

	"hypergrid-ssn/x/hypergridssn/types"

	solana "hypergrid-ssn/tools"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateHypergridNode(goCtx context.Context, msg *types.MsgCreateHypergridNode) (*types.MsgCreateHypergridNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetHypergridNode(
		ctx,
		msg.Pubkey,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	if msg.Role == 3 { //grid node
		sonic_grid_rpc := ""
		nodes := k.GetAllHypergridNode(goCtx)
		for _, node := range nodes {
			// get sonic grid node from the list
			if node.Role == 2 {
				sonic_grid_rpc = node.Rpc
				break
			}
		}

		if sonic_grid_rpc == "" {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "sonic grid rpc not found")
		}

		sig, err := solana.InitializeDataAccount(sonic_grid_rpc, msg.Pubkey, msg.DataAccount, uint32(msg.Role))
		if err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, err.Error())
		}

		println("signature: ", sig)
	}

	var hypergridNode = types.HypergridNode{
		Creator:     msg.Creator,
		Pubkey:      msg.Pubkey,
		Name:        msg.Name,
		Rpc:         msg.Rpc,
		DataAccount: msg.DataAccount,
		Role:        msg.Role,
		Starttime:   msg.Starttime,
	}

	k.SetHypergridNode(
		ctx,
		hypergridNode,
	)
	return &types.MsgCreateHypergridNodeResponse{}, nil
}

func (k msgServer) UpdateHypergridNode(goCtx context.Context, msg *types.MsgUpdateHypergridNode) (*types.MsgUpdateHypergridNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetHypergridNode(
		ctx,
		msg.Pubkey,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var hypergridNode = types.HypergridNode{
		Creator:     msg.Creator,
		Pubkey:      msg.Pubkey,
		Name:        msg.Name,
		Rpc:         msg.Rpc,
		DataAccount: msg.DataAccount,
		Role:        msg.Role,
		Starttime:   msg.Starttime,
	}

	k.SetHypergridNode(ctx, hypergridNode)

	return &types.MsgUpdateHypergridNodeResponse{}, nil
}

func (k msgServer) DeleteHypergridNode(goCtx context.Context, msg *types.MsgDeleteHypergridNode) (*types.MsgDeleteHypergridNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetHypergridNode(
		ctx,
		msg.Pubkey,
	)
	// isFound := k.HasHypergridNode(ctx, msg.Pubkey)

	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHypergridNode(
		ctx,
		msg.Pubkey,
	)

	return &types.MsgDeleteHypergridNodeResponse{}, nil
}
