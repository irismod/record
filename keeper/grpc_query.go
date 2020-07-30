package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irismod/record/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Record(c context.Context, req *types.QueryRecordRequest) (*types.QueryRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	record, _ := k.GetRecord(ctx, req.Recordid)
	return &types.QueryRecordResponse{
		Record: &record,
	}, nil
}
