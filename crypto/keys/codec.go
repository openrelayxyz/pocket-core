package keys

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/crypto"
)

// LegacyCodec defines the codec required for keys and info
var LegacyCodec *codec.LegacyAmino

var cdc = codec.NewProtoCodec(types.NewInterfaceRegistry())

func init() {
	LegacyCodec = codec.NewLegacyAminoCodec()
	crypto.RegisterCrypto(LegacyCodec, nil)
	RegisterCodec(LegacyCodec)
	LegacyCodec.Seal()
}

// RegisterCodec registers concrete types and interfaces on the given codec.
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(KeyPair{}, "crypto/keys/keypair", nil)
}
