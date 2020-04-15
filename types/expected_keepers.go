package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/exported"
)

// AccountKeeper defines the expected account keeper for query account
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) auth.Account
}
