package types

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
	types3 "github.com/pokt-network/pocket-core/x/apps/types"
	"github.com/pokt-network/pocket-core/x/auth/exported"
	types5 "github.com/pokt-network/pocket-core/x/gov/types"
	types2 "github.com/pokt-network/pocket-core/x/nodes/types"
	types4 "github.com/pokt-network/pocket-core/x/pocketcore/types"
)

// RegisterCodec registers concrete types on the codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface("x.auth.ModuleAccount", (*exported.ModuleAccountI)(nil), &ModuleAccountEncodable{})
	cdc.RegisterInterface("x.auth.Account", (*exported.Account)(nil), &BaseAccountEncodable{}, &ModuleAccountEncodable{})
	cdc.RegisterInterface("x.auth.Supply", (*exported.SupplyI)(nil), &Supply{})
	cdc.RegisterStructure(&BaseAccount{}, "posmint/Account")
	cdc.RegisterStructure(LegacyStdTx{}, "posmint/StdTx")
	cdc.RegisterStructure(&Supply{}, "posmint/Supply")
	cdc.RegisterStructure(&ModuleAccount{}, "posmint/ModuleAccount")
	cdc.RegisterImplementation((*sdk.Tx)(nil), &StdTx{})
	cdc.RegisterImplementation((*sdk.Msg)(nil), &types2.MsgNodeStake{}, &types2.MsgUnjail{}, &types2.MsgBeginUnstake{},
		&types2.MsgSend{}, types2.MsgStake{}, &types3.MsgApplicationStake{}, &types3.MsgBeginAppUnstake{}, &types3.MsgAppUnjail{}, types3.MsgAppStake{},
		&types4.MsgClaim{}, &types4.MsgProtoProof{}, types4.MsgProof{}, &types5.MsgChangeParam{}, &types5.MsgDAOTransfer{}, &types5.MsgUpgrade{})
	cdc.RegisterImplementation((*sdk.LegacyMsg)(nil), &types2.MsgNodeStake{}, &types2.MsgUnjail{}, &types2.MsgBeginUnstake{},
		&types2.MsgSend{}, types2.MsgStake{}, &types3.MsgApplicationStake{}, &types3.MsgBeginAppUnstake{}, &types3.MsgAppUnjail{}, types3.MsgAppStake{},
		&types4.MsgClaim{}, &types4.MsgProtoProof{}, types4.MsgProof{}, &types5.MsgChangeParam{}, &types5.MsgDAOTransfer{}, &types5.MsgUpgrade{})
}

// module wide codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.NewCodec(types.NewInterfaceRegistry())
	RegisterCodec(ModuleCdc)
	crypto.RegisterAmino(ModuleCdc.AminoCodec().Amino)
}
