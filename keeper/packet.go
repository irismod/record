package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctypes "github.com/cosmos/cosmos-sdk/x/ibc/types"
	"github.com/irismod/record/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

// BindPort defines a wrapper function for the ort Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	// Set the portID into our store so we can retrieve it later
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.PortKey), []byte(portID))

	cap := k.portKeeper.BindPort(ctx, portID)
	return k.scopedKeeper.ClaimCapability(ctx, cap, ibctypes.PortPath(portID))
}

// GetPort returns the portID for the record module.
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get([]byte(types.PortKey)))
}

func (k Keeper) PacketExecute(ctx sdk.Context, packet types.Packet) ([]byte, error) {
	content := types.Content{
		Digest:     packet.Digest,
		DigestAlgo: packet.DigestAlgo,
		URI:        packet.URI,
		Meta:       packet.Metadata,
	}

	var pubKey crypto.PubKey
	if err := k.cdc.UnmarshalJSON(packet.PubKey, &pubKey); err != nil {
		return nil, err
	}
	creator := sdk.AccAddress(pubKey.Address().Bytes())
	record := types.NewRecord(tmhash.Sum(ctx.TxBytes()), []types.Content{content}, creator)
	recordID := k.AddRecord(ctx, record)
	return recordID, nil
}
