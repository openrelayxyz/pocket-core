package types

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
)

// RegisterCodec registers concrete types on the codec
func RegisterCodec(amino *codec.LegacyAmino, proto *codec.ProtoCodec) {
	amino.RegisterConcrete(MsgAppStake{}, "apps/MsgAppStake", nil)
	amino.RegisterConcrete(MsgBeginAppUnstake{}, "apps/MsgAppBeginUnstake", nil)
	amino.RegisterConcrete(MsgAppUnjail{}, "apps/MsgAppUnjail", nil)

	proto.RegisterImplementation((*sdk.Msg)(nil), &MsgAppStake{}, &MsgBeginAppUnstake{}, &MsgAppUnjail{})
	ModuleCdc = proto
}

// module wide codec
var ModuleCdc *codec.ProtoCodec
var LegacyModuleCdc *codec.LegacyAmino

func init() {
	ModuleCdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
	LegacyModuleCdc = codec.NewLegacyAminoCodec()
	RegisterCodec(LegacyModuleCdc, ModuleCdc)
	crypto.RegisterCrypto(LegacyModuleCdc, ModuleCdc)
	LegacyModuleCdc.Seal()
}
