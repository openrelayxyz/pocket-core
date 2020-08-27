package keeper

import (
	"fmt"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/auth/exported"
	"github.com/pokt-network/pocket-core/x/auth/types"
)

// GetSupply retrieves the Supply from store
func (k Keeper) GetSupply(ctx sdk.Ctx) (supply exported.SupplyI) {
	store := ctx.KVStore(k.storeKey)
	b, _ := store.Get(types.SupplyKeyPrefix)
	if b == nil {
		ctx.Logger().Error(fmt.Sprintf("stored supply should not have been nil, at height: %d", ctx.BlockHeight()))
		return
	}
	var s types.Supply
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &s)
	return &s
}

// SetSupply sets the Supply to store
func (k Keeper) SetSupply(ctx sdk.Ctx, supply exported.SupplyI) {
	store := ctx.KVStore(k.storeKey)
	s := supply.(*types.Supply)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(s)
	_ = store.Set(types.SupplyKeyPrefix, b)
}
