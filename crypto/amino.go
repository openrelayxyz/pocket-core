package crypto

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

var cdc *codec.LegacyAmino

func init() {
	cdc = codec.NewLegacyAminoCodec()
	RegisterCrypto(cdc, nil)
}

// RegisterCrypto registers all go-crypto related types in the given (cdc) codec.
func RegisterCrypto(cdc *codec.LegacyAmino, _ *codec.ProtoCodec) {
	cdc.RegisterInterface((*PublicKey)(nil), nil)
	cdc.RegisterInterface((*PrivateKey)(nil), nil)
	cdc.RegisterConcrete(Ed25519PublicKey{}, "crypto/ed25519_public_key", nil)
	cdc.RegisterConcrete(Ed25519PrivateKey{}, "crypto/ed25519_private_key", nil)
	cdc.RegisterConcrete(Secp256k1PublicKey{}, "crypto/secp256k1_public_key", nil)
	cdc.RegisterConcrete(Secp256k1PrivateKey{}, "crypto/secp256k1_private_key", nil)
	cdc.RegisterInterface((*MultiSig)(nil), nil)
	cdc.RegisterInterface((*PublicKeyMultiSig)(nil), nil)
	cdc.RegisterConcrete(PublicKeyMultiSignature{}, "crypto/public_key_multi_signature", nil)
	cdc.RegisterConcrete(MultiSignature{}, "crypto/multi_signature", nil)
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PubKeyEd25519{}, ed25519.PubKeyAminoName, nil)
	cdc.RegisterConcrete(secp256k1.PubKeySecp256k1{}, secp256k1.PubKeyAminoName, nil)
	cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PrivKeyEd25519{}, ed25519.PrivKeyAminoName, nil)
	cdc.RegisterConcrete(secp256k1.PrivKeySecp256k1{}, secp256k1.PrivKeyAminoName, nil)
}
