package types

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/crypto"
	"github.com/pokt-network/pocket-core/x/nodes/exported"
	nodesTypes "github.com/pokt-network/pocket-core/x/nodes/types"
)

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

// RegisterCodec registers concrete types on the codec
func RegisterCodec(amino *codec.LegacyAmino, proto *codec.ProtoCodec) {
	amino.RegisterConcrete(MsgClaim{}, "pocketcore/claim", nil)
	amino.RegisterConcrete(MsgProof{}, "pocketcore/proof", nil)
	amino.RegisterConcrete(Relay{}, "pocketcore/relay", nil)
	amino.RegisterConcrete(Session{}, "pocketcore/session", nil)
	amino.RegisterConcrete(RelayResponse{}, "pocketcore/relay_response", nil)
	amino.RegisterConcrete(RelayProof{}, "pocketcore/relay_proof", nil)
	amino.RegisterConcrete(ChallengeProofInvalidData{}, "pocketcore/challenge_proof_invalid_data", nil)
	amino.RegisterConcrete(EvidenceEncodable{}, "pocketcore/evidence_persisted", nil)
	amino.RegisterConcrete(nodesTypes.Validator{}, "pos/Validator", nil) // todo does this really need to depend on nodes/types
	amino.RegisterInterface((*Proof)(nil), nil)
	amino.RegisterInterface((*exported.ValidatorI)(nil), nil)
	proto.Register("x.pocketcore.Proof", (*Proof)(nil), &RelayProof{}, &ChallengeProofInvalidData{})
}
