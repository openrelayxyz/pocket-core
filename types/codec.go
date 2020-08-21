package types

import "github.com/pokt-network/pocket-core/codec"

// Register the sdk message type
func RegisterCodec(amino *codec.LegacyAmino, proto *codec.ProtoCodec) {
	amino.RegisterInterface((*Msg)(nil), nil)
	amino.RegisterInterface((*Tx)(nil), nil)
	proto.Register("types/msg", (*Msg)(nil))
	proto.Register("types/tx", (*Tx)(nil))
}
