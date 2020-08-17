package types

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/crypto"
)

// Register concrete types on codec
func RegisterCodec(cdc *codec.ProtoCodec) {
	//cdc.RegisterConcrete(MsgStake{}, "pos/MsgStake", nil)
	//cdc.RegisterConcrete(MsgBeginUnstake{}, "pos/MsgBeginUnstake", nil)
	//cdc.RegisterConcrete(MsgUnjail{}, "pos/MsgUnjail", nil)
	//cdc.RegisterConcrete(MsgSend{}, "pos/Send", nil)
}

var ModuleCdc *codec.ProtoCodec       // generic sealed codec to be used throughout this module
var ModuleAminoCdc *codec.LegacyAmino // generic sealed codec to be used throughout this module

func init() {

	ModuleCdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
	ModuleAminoCdc = codec.New()
	RegisterCodec(ModuleCdc)
	crypto.RegisterAmino(ModuleAminoCdc)
	ModuleAminoCdc.Seal()
}
