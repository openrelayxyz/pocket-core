package types

import (
	"github.com/pokt-network/pocket-core/codec"
)

// Register concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAppStake{}, "apps/MsgAppStake", nil)
	cdc.RegisterConcrete(MsgBeginAppUnstake{}, "apps/MsgAppBeginUnstake", nil)
	cdc.RegisterConcrete(MsgAppUnjail{}, "apps/MsgAppUnjail", nil)
}

var ModuleCdc *codec.ProtoCodec // generic sealed codec to be used throughout this module

// RegisterInterface associates protoName with AccountI interface
// and creates a registry of it's concrete implementations
// NOTE: applications has no interface that requries to be registeredj
func RegisterInterfaces(registry types.InterfaceRegistry) {
}


func init() {
	ModuleCdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal())
}
