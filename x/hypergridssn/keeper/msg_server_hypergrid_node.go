package keeper

import (
	"context"

	"hypergrid-ssn/x/hypergridssn/types"

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

	var hypergridNode = types.HypergridNode{
		Creator:   msg.Creator,
		Pubkey:    msg.Pubkey,
		Name:      msg.Name,
		Rpc:       msg.Rpc,
		Role:      msg.Role,
		Starttime: msg.Starttime,
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
		Creator:   msg.Creator,
		Pubkey:    msg.Pubkey,
		Name:      msg.Name,
		Rpc:       msg.Rpc,
		Role:      msg.Role,
		Starttime: msg.Starttime,
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
