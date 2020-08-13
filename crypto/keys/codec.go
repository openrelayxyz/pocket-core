package keys

import (

	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/crypto"
)

// CryptoCdc defines the codec required for keys and info
var CryptoCdc *codec.LegacyAmino

func init() {
	CryptoCdc = codec.New()
	crypto.RegisterAmino(CryptoCdc)
	RegisterCodec(CryptoCdc)
	CryptoCdc.Seal()
}

// RegisterCodec registers concrete types and interfaces on the given codec.
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(KeyPair{}, "crypto/keys/keypair", nil)
}
