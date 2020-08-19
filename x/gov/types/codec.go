package types

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
)

// RegisterCodec registers concrete types on the codec
func RegisterCodec(amino *codec.LegacyAmino, proto *codec.ProtoCodec) {
	amino.RegisterConcrete(MsgChangeParam{}, "gov/msg_change_param", nil)
	amino.RegisterConcrete(MsgDAOTransfer{}, "gov/msg_dao_transfer", nil)
	amino.RegisterInterface((*interface{})(nil), nil)
	amino.RegisterConcrete(ACL{}, "gov/non_map_acl", nil)
	amino.RegisterConcrete(Upgrade{}, "gov/upgrade", nil)
	amino.RegisterConcrete(MsgUpgrade{}, "gov/msg_upgrade", nil)

	proto.RegisterImplementation((*sdk.Msg)(nil), &MsgChangeParam{}, &MsgDAOTransfer{})
	ModuleCdc = proto
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
