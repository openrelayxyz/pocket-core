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
	supply, err := k.DecodeSupply(b)
	if err != nil {
		ctx.Logger().Error(fmt.Sprintf(err.Error()))
		return
	}
	return supply
}

// SetSupply sets the Supply to store
func (k Keeper) SetSupply(ctx sdk.Ctx, supply exported.SupplyI) {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.EncodeSupply(supply)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}
	_ = store.Set(types.SupplyKeyPrefix, bz)
}

func (k Keeper) EncodeSupply(supply exported.SupplyI) ([]byte, error) {
	var bz []byte
	var err error
	if BLOCKHEIGHTPASSED {
		bz, err = k.cdc.ProtoMarshalBinaryLengthPrefixed(supply.(*types.Supply))
	} else {
		bz, err = k.cdc.LegacyMarshalBinaryLengthPrefixed(supply) // TODO only kept this way for backwards compatibility.. test if breaks when using supply.(*Supply)
	}
	return bz, err
}

func (k Keeper) DecodeSupply(bz []byte) (exported.SupplyI, error) {
	var supply types.Supply
	err := k.cdc.UnmarshalBinaryLengthPrefixed(bz, &supply)
	return supply, err
}
