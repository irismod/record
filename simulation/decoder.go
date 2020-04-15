package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	tmkv "github.com/tendermint/tendermint/libs/kv"

	"github.com/irismod/record/types"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding slashing type
func DecodeStore(cdc *codec.Codec, kvA, kvB tmkv.Pair) string {
	switch {
	case bytes.Equal(kvA.Key[:1], types.RecordKey):
		var recordA, recordB types.Record
		cdc.MustUnmarshalBinaryLengthPrefixed(kvA.Value, &recordA)
		cdc.MustUnmarshalBinaryLengthPrefixed(kvB.Value, &recordB)
		return fmt.Sprintf("%v\n%v", recordA, recordB)
	default:
		panic(fmt.Sprintf("invalid record key prefix %X", kvA.Key[:1]))
	}
}
