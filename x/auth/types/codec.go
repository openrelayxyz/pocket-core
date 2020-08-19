package types

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/auth/exported"
)

// RegisterCodec registers concrete types on the codec
func RegisterCodec(amino *codec.LegacyAmino, proto *codec.ProtoCodec) {
	amino.RegisterInterface((*exported.Account)(nil), nil)
	amino.RegisterConcrete(&BaseAccount{}, "posmint/Account", nil)
	amino.RegisterConcrete(StdTx{}, "posmint/StdTx", nil)
	amino.RegisterInterface((*exported.ModuleAccountI)(nil), nil)
	amino.RegisterInterface((*exported.SupplyI)(nil), nil)
	amino.RegisterConcrete(&ModuleAccount{}, "posmint/ModuleAccount", nil)
	amino.RegisterConcrete(&Supply{}, "posmint/Supply", nil)
	proto.Register("x.auth.Account", (*exported.Account)(nil), &BaseAccountEncodable{}, &ModuleAccountEncodable{})
	proto.Register("x.auth.ModuleAccount", (*exported.Account)(nil), &ModuleAccountEncodable{})
	proto.Register("x.auth.SupplyI", (*exported.SupplyI)(nil), &Supply{})
	proto.RegisterImplementation((*sdk.Tx)(nil), &StdTx{})
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
