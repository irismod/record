package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/bytes"
)

// NewRecord constructs a record
func NewRecord(txHash bytes.HexBytes, contents []Content, creator sdk.AccAddress) Record {
	return Record{
		TxHash:   txHash,
		Contents: contents,
		Creator:  creator,
	}
}
