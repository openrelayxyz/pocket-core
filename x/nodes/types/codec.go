package types

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/crypto"

	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/nodes/exported"
)

// Register concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterStructure(MsgStake{}, "pos/MsgStake")
	cdc.RegisterStructure(MsgBeginUnstake{}, "pos/MsgBeginUnstake")
	cdc.RegisterStructure(MsgUnjail{}, "pos/MsgUnjail")
	cdc.RegisterStructure(MsgSend{}, "pos/Send")
	cdc.RegisterStructure(LegacyMsgStake{}, "pos/LegacyMsgStake")
	cdc.RegisterImplementation((*sdk.Msg)(nil), &MsgStake{}, &MsgUnjail{}, &MsgBeginUnstake{}, &MsgSend{}, LegacyMsgStake{})
	cdc.RegisterImplementation((*sdk.LegacyMsg)(nil), &MsgStake{}, &MsgUnjail{}, &MsgBeginUnstake{}, &MsgSend{}, LegacyMsgStake{})
	cdc.RegisterInterface("nodes/validatorI", (*exported.ValidatorI)(nil), &ValidatorProto{})

}

var ModuleCdc *codec.Codec // generic sealed codec to be used throughout this module

func init() {
	ModuleCdc = codec.NewCodec(types.NewInterfaceRegistry())
	RegisterCodec(ModuleCdc)
	crypto.RegisterAmino(ModuleCdc.AminoCodec().Amino)
	ModuleCdc.AminoCodec().Seal()
}
