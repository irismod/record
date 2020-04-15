package simulation

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/irismod/record/types"
)

// RandomizedGenState generates a random GenesisState for record
func RandomizedGenState(simState *module.SimulationState) {
	records := make([]types.Record, simState.Rand.Intn(100))

	for i := 0; i < len(records); i++ {
		records[i], _ = genRecord(simState.Rand, simState.Accounts)
	}

	recordGenesis := types.NewGenesisState(records)

	fmt.Printf("Selected randomly generated record parameters:\n%s\n", codec.MustMarshalJSONIndent(simState.Cdc, recordGenesis))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(recordGenesis)
}
