package keeper

// DONTCOVER

import (
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/crypto"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/store"
	sdk "github.com/pokt-network/pocket-core/types"
	authTypes "github.com/pokt-network/pocket-core/x/auth/types"
	govKeeper "github.com/pokt-network/pocket-core/x/gov/keeper"
	govTypes "github.com/pokt-network/pocket-core/x/gov/types"
)

type testInput struct {
	cdc    *codec.ProtoCodec
	ctx    sdk.Context
	Keeper Keeper
}

func setupTestInput() testInput {
	db := dbm.NewMemDB()

	cdc := codec.NewLegacyAminoCodec()
	proto := codec.NewProtoCodec(types.NewInterfaceRegistry())
	authTypes.RegisterCodec(cdc, proto)
	crypto.RegisterCrypto(cdc, nil)

	authCapKey := sdk.NewKVStoreKey("auth")
	keyParams := sdk.ParamsKey
	tkeyParams := sdk.ParamsTKey

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(authCapKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)
	_ = ms.LoadLatestVersion()
	akSubspace := sdk.NewSubspace(authTypes.DefaultCodespace)
	ak := NewKeeper(
		cdc, proto, authCapKey, akSubspace, nil,
	)
	govKeeper.NewKeeper(cdc, proto, sdk.ParamsKey, sdk.ParamsTKey, govTypes.DefaultCodespace, ak, akSubspace)
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	ak.SetParams(ctx, authTypes.DefaultParams())
	return testInput{Keeper: ak, cdc: proto, ctx: ctx}
}
