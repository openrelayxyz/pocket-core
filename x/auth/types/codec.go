package types

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/x/auth/exported"
)

// RegisterCodec registers concrete types on the codec
func RegisterCodec(cdc *codec.ProtoCodec) {
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&BaseAccount{}, "posmint/Account", nil)
	cdc.RegisterConcrete(StdTx{}, "posmint/StdTx", nil)
	cdc.RegisterInterface((*exported.ModuleAccountI)(nil), nil)
	cdc.RegisterInterface((*exported.SupplyI)(nil), nil)
	cdc.RegisterConcrete(&ModuleAccount{}, "posmint/ModuleAccount", nil)
	cdc.RegisterConcrete(&Supply{}, "posmint/Supply", nil)
}


// RegisterInterface associates protoName with AccountI interface
// and creates a registry of it's concrete implementations
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"pocket.auth.Account",
		(*exported.Account)(nil),
		&BaseAccount{},
		&ModuleAccount{},
	)

	registry.RegisterInterface(
		"pocket.auth.ModuleAccountI",
		(*exported.ModuleAccountI)(nil),
		&ModuleAccount{},
	)

	registry.RegisterInterface(
		"pocket.auth.SupplyI",
		(*exported.SupplyI)(nil),
		&Supply{},
	)
}

// module wide codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
