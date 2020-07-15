package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"

	"github.com/irismod/record/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

func (k Keeper) PacketExecute(ctx sdk.Context, packet channeltypes.Packet) ([]byte, error) {
	var data types.Packet
	if err := k.cdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal record packet data: %s", err.Error())
	}

	content := types.Content{
		Digest:     data.Digest,
		DigestAlgo: data.DigestAlgo,
		URI:        data.URI,
		Meta:       data.Metadata,
	}

	var pubKey crypto.PubKey
	if err := k.cdc.UnmarshalJSON(data.PubKey, &pubKey); err != nil {
		return nil, err
	}
	creator := sdk.AccAddress(pubKey.Address().Bytes())
	record := types.NewRecord(tmhash.Sum(ctx.TxBytes()), []types.Content{content}, creator)
	recordID := k.AddRecord(ctx, record)
	return recordID, nil
}
