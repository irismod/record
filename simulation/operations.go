package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/irismod/record/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateRecord = "op_weight_msg_create_record"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simulation.AppParams,
	cdc *codec.Codec,
	ak types.AccountKeeper) simulation.WeightedOperations {
	var weightCreate int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreateRecord, &weightCreate, nil,
		func(_ *rand.Rand) {
			weightCreate = 50
		},
	)
	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightCreate,
			SimulateCreateRecord(ak),
		),
	}
}

// SimulateCreateRecord tests and runs a single msg create a new record
func SimulateCreateRecord(ak types.AccountKeeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		record, err := genRecord(r, accs)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		msg := types.NewMsgCreateRecord(record.Contents, record.Creator)

		simAccount, found := simulation.FindAccount(accs, record.Creator)
		if !found {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("account %s not found", record.Creator)
		}
		account := ak.GetAccount(ctx, msg.Creator)

		spendable := account.SpendableCoins(ctx.BlockTime())
		fees, err := simulation.RandomFees(r, ctx, spendable)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err = app.Deliver(tx); err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, "simulate issue token"), nil, nil
	}
}

func genRecord(r *rand.Rand, accs []simulation.Account) (types.Record, error) {
	var record types.Record
	txHash := make([]byte, 32)
	_, err := r.Read(txHash)
	if err != nil {
		return record, err
	}

	record.TxHash = txHash

	for i := 0; i <= r.Intn(10); i++ {
		record.Contents = append(record.Contents, types.Content{
			Digest:     "test",
			DigestAlgo: "SHA256",
		})
	}

	acc, _ := simulation.RandomAcc(r, accs)
	record.Creator = acc.Address

	return record, nil
}
