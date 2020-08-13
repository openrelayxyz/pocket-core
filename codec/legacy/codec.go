package legacy

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/crypto"
)

// Cdc defines a global generic sealed Amino codec to be used throughout sdk. It
// has all Tendermint crypto and evidence types registered.
//
// TODO: Deprecated - remove this global.
var Cdc *codec.LegacyAmino

func init() {
	Cdc = codec.New()
	crypto.RegisterAmino(Cdc)
	codec.RegisterEvidences(Cdc)
	Cdc.Seal()
}
