package keeper

import (
	"context"
	solana "hypergrid-ssn/tools"
	"hypergrid-ssn/x/hypergridssn/types"
	"strconv"

	cmtjson "github.com/cometbft/cometbft/libs/json"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSolanaAccount(goCtx context.Context, msg *types.MsgCreateSolanaAccount) (*types.MsgCreateSolanaAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetSolanaAccount(
		ctx,
		msg.Address,
		msg.Version,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	//get hypergrid node from msg.Source
	node, found := k.GetHypergridNode(goCtx, msg.Source)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "source not found")
	}

	// get account info from solana
	// resp, err := solana.GetAccountInfo(node.Rpc, msg.Address)
	resp, err := solana.GetAccountFromOracle(node.Rpc, msg.Address, msg.Version)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, err.Error())
	}

	value, err := cmtjson.Marshal(resp.Value)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, err.Error())
	}

	var solanaAccount = types.SolanaAccount{
		Creator: msg.Creator,
		Address: msg.Address,
		Version: msg.Version,
		Source:  node.Pubkey,
		Slot:    strconv.FormatUint(resp.RPCContext.Context.Slot, 10),
		Value:   string(value),
	}

	k.SetSolanaAccount(
		ctx,
		solanaAccount,
	)
	return &types.MsgCreateSolanaAccountResponse{
		// SolanaAccount: solanaAccount,
	}, nil
}

func (k msgServer) UpdateSolanaAccount(goCtx context.Context, msg *types.MsgUpdateSolanaAccount) (*types.MsgUpdateSolanaAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSolanaAccount(
		ctx,
		msg.Address,
		msg.Version,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	//get hypergrid node from msg.Pubkey
	node, found := k.GetHypergridNode(goCtx, valFound.Source)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "node not found")
	}

	// get account info from solana
	// resp, err := solana.GetAccountInfo(node.Rpc, msg.Address)
	resp, err := solana.GetAccountFromOracle(node.Rpc, msg.Address, msg.Version)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, err.Error())
	}

	value, err := cmtjson.Marshal(resp.Value)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, err.Error())
	}

	var solanaAccount = types.SolanaAccount{
		Creator: msg.Creator,
		Address: msg.Address,
		Version: msg.Version,
		Source:  valFound.Source,
		Slot:    strconv.FormatUint(resp.RPCContext.Context.Slot, 10),
		Value:   string(value),
	}

	k.SetSolanaAccount(ctx, solanaAccount)

	return &types.MsgUpdateSolanaAccountResponse{
		// SolanaAccount: solanaAccount,
	}, nil
}

func (k msgServer) DeleteSolanaAccount(goCtx context.Context, msg *types.MsgDeleteSolanaAccount) (*types.MsgDeleteSolanaAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSolanaAccount(
		ctx,
		msg.Address,
		msg.Version,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSolanaAccount(
		ctx,
		msg.Address,
		msg.Version,
	)

	return &types.MsgDeleteSolanaAccountResponse{}, nil
}
