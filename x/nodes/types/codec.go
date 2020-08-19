package types

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/nodes/exported"
)

// RegisterCodec registers concrete types on the codec
func RegisterCodec(amino *codec.LegacyAmino, proto *codec.ProtoCodec) {
	amino.RegisterConcrete(MsgStake{}, "pos/MsgStake", nil)
	amino.RegisterConcrete(MsgBeginUnstake{}, "pos/MsgBeginUnstake", nil)
	amino.RegisterConcrete(MsgUnjail{}, "pos/MsgUnjail", nil)
	amino.RegisterConcrete(MsgSend{}, "pos/Send", nil)

	proto.RegisterImplementation((*sdk.Msg)(nil), &MsgStake{}, &MsgUnjail{}, &MsgBeginUnstake{}, &MsgSend{})
	amino.RegisterInterface((*exported.ValidatorI)(nil), nil)
	proto.Register("nodes/validatorI", (*exported.ValidatorI)(nil), &ValidatorProto{})
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
