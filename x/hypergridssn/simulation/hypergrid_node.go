package simulation

import (
	"math/rand"
	"strconv"

	"hypergrid-ssn/x/hypergridssn/keeper"
	"hypergrid-ssn/x/hypergridssn/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateHypergridNode(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateHypergridNode{
			Creator: simAccount.Address.String(),
			Pubkey:  strconv.Itoa(i),
		}

		_, found := k.GetHypergridNode(ctx, msg.Pubkey)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "HypergridNode already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           moduletestutil.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateHypergridNode(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount       = simtypes.Account{}
			hypergridNode    = types.HypergridNode{}
			msg              = &types.MsgUpdateHypergridNode{}
			allHypergridNode = k.GetAllHypergridNode(ctx)
			found            = false
		)
		for _, obj := range allHypergridNode {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				hypergridNode = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "hypergridNode creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Pubkey = hypergridNode.Pubkey

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           moduletestutil.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteHypergridNode(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount       = simtypes.Account{}
			hypergridNode    = types.HypergridNode{}
			msg              = &types.MsgUpdateHypergridNode{}
			allHypergridNode = k.GetAllHypergridNode(ctx)
			found            = false
		)
		for _, obj := range allHypergridNode {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				hypergridNode = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "hypergridNode creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Pubkey = hypergridNode.Pubkey

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           moduletestutil.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
