package keys

import (
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/crypto"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"

	"github.com/pokt-network/pocket-core/codec"
)

var cdc *codec.Codec

func init() {
	cdc = codec.NewCodec(types.NewInterfaceRegistry())
	cryptoAmino.RegisterAmino(cdc.AminoCodec().Amino)
	crypto.RegisterAmino(cdc.AminoCodec().Amino)
	cdc.AminoCodec().RegisterConcrete(KeyPair{}, "crypto/keys/keypair", nil)
	cdc.AminoCodec().Seal()
}
