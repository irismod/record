package simulation

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto/secp256k1"
	tmkv "github.com/tendermint/tendermint/libs/kv"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irismod/record/types"
)

var (
	creatorPk1   = secp256k1.GenPrivKey().PubKey()
	creatorAddr1 = sdk.AccAddress(creatorPk1.Address())
)

func makeTestCodec() (codec.Marshaler, *codec.Codec) {
	cdc := std.MakeCodec(simapp.ModuleBasics)
	types.RegisterCodec(cdc)
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	sdk.RegisterInterfaces(interfaceRegistry)
	simapp.ModuleBasics.RegisterInterfaceModules(interfaceRegistry)
	encodingConfig := simapp.MakeEncodingConfig()
	appCodec := encodingConfig.Marshaler
	return appCodec, cdc
}

func TestDecodeStore(t *testing.T) {
	cdc, _ := makeTestCodec()
	dec := NewDecodeStore(cdc)

	txHash := make([]byte, 32)
	_, _ = rand.Read(txHash)
	record := types.NewRecord(txHash, nil, creatorAddr1)

	kvPairs := tmkv.Pairs{
		tmkv.Pair{Key: types.GetRecordKey(txHash), Value: cdc.MustMarshalBinaryBare(&record)},
		tmkv.Pair{Key: []byte{0x99}, Value: []byte{0x99}},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Record", fmt.Sprintf("%v\n%v", record, record)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}
