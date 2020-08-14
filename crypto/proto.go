package crypto

import (
	"github.com/pokt-network/pocket-core/codec"
	types "github.com/pokt-network/pocket-core/codec/types"
)

var proto *codec.ProtoCodec

func init() {
	proto = codec.NewProtoCodec(RegisterProto())
}

// Registerproto registers all go-crypto related  (proto) codec.
func RegisterProto() types.InterfaceRegistry {
	registry := types.NewInterfaceRegistry()
	registry.RegisterInterface("PublicKey", (*PublicKey)(nil), &Ed25519PublicKey{}, &Secp256k1PublicKey{})
	registry.RegisterInterface("PrivateKey", (*PrivateKey)(nil), &Ed25519PrivateKey{}, &Secp256k1PrivateKey{})
	registry.RegisterInterface("MultiSig", (*MultiSig)(nil), &MultiSignature{})
	registry.RegisterInterface("PublicKeyMultiSig", (*PublicKeyMultiSig)(nil), &PublicKeyMultiSignature{}, nil)
	return registry
}
