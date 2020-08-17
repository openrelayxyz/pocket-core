package types

import (
	"fmt"
	"github.com/pokt-network/pocket-core/types"
	"github.com/willf/bloom"
	"strings"
)

// "Evidence" - A proof of work/burn for nodes.
type Evidence struct {
	Bloom         bloom.BloomFilter        `json:"bloom_filter"` // used to check if proof contains
	SessionHeader `json:"evidence_header"` // the session h serves as an identifier for the EvidenceEncodable
	NumOfProofs   int64                    `json:"num_of_proofs"` // the total number of proofs in the EvidenceEncodable
	Proofs        []Proof                  `json:"proofs"`        // a slice of Proof objects (Proof per relay or challenge)
	EvidenceType  EvidenceType             `json:"evidence_type"`
}

// "GenerateMerkleRoot" - Generates the merkle root for an EvidenceEncodable object
func (e *Evidence) GenerateMerkleRoot() (root HashRange) {
	// generate the root object
	root, sortedProofs := GenerateRoot(e.Proofs)
	// sort the proofs
	e.Proofs = sortedProofs
	// set the EvidenceEncodable in cache
	SetEvidence(*e)
	return
}

// "AddProof" - Adds a proof obj to the EvidenceEncodable field
func (e *Evidence) AddProof(p Proof) {
	// add proof to EvidenceEncodable
	e.Proofs = append(e.Proofs, p)
	// increment total proof count
	e.NumOfProofs = e.NumOfProofs + 1
	// add proof to bloom filter
	e.Bloom.Add(p.Hash())
}

// "GenerateMerkleProof" - Generates the merkle Proof for an EvidenceEncodable
func (e *Evidence) GenerateMerkleProof(index int) (proof MerkleProof, leaf Proof) {
	// generate the merkle proof
	proof, leaf = GenerateProofs(e.Proofs, index)
	// set the EvidenceEncodable in memory
	SetEvidence(*e)
	return
}

// "Evidence" - A proof of work/burn for nodes.
//type EvidenceEncodable struct {
//	BloomBytes    []byte                   `json:"bloom_bytes"`
//	SessionHeader `json:"evidence_header"` // the session h serves as an identifier for the EvidenceEncodable
//	NumOfProofs   int64                    `json:"num_of_proofs"` // the total number of proofs in the EvidenceEncodable
//	Proofs        []Proof                  `json:"proofs"`        // a slice of Proof objects (Proof per relay or challenge)
//	EvidenceType  EvidenceType             `json:"evidence_type"`
//}

var _ CacheObject = Evidence{} // satisfies the cache object interface

func (e Evidence) Marshal() ([]byte, error) {
	encodedBloom, err := e.Bloom.GobEncode()
	if err != nil {
		return nil, err
	}
	ep := EvidenceEncodable{
		BloomBytes:    encodedBloom,
		SessionHeader: &e.SessionHeader,
		NumOfProofs:   e.NumOfProofs,
		Proofs:        e.Proofs,
		EvidenceType:  e.EvidenceType,
	}
	return ModuleCdc.MarshalBinaryBare(ep)
}

func (e Evidence) Unmarshal(b []byte) (CacheObject, error) {
	ep := EvidenceEncodable{}
	err := ModuleCdc.UnmarshalBinaryBare(b, &ep)
	if err != nil {
		return Evidence{}, fmt.Errorf("could not unmarshal into EvidenceEncodable from cache, moduleCdc unmarshal binary bare: %s", err.Error())
	}
	bloomFilter := bloom.BloomFilter{}
	err = bloomFilter.GobDecode(ep.BloomBytes)
	if err != nil {
		return Evidence{}, fmt.Errorf("could not unmarshal into EvidenceEncodable from cache, bloom bytes gob decode: %s", err.Error())
	}
	evidence := Evidence{
		Bloom:         bloomFilter,
		SessionHeader: *ep.SessionHeader,
		NumOfProofs:   ep.NumOfProofs,
		Proofs:        ep.Proofs,
		EvidenceType:  ep.EvidenceType}
	return evidence, nil
}

func (e Evidence) Key() ([]byte, error) {
	return KeyForEvidence(e.SessionHeader, e.EvidenceType)
}

// "EvidenceType" type to distinguish the types of EvidenceEncodable (relay/challenge)
type EvidenceType int

const (
	RelayEvidence EvidenceType = iota + 1 // essentially an enum for EvidenceEncodable types
	ChallengeEvidence
)

// "Convert EvidenceEncodable type to bytes
func (et EvidenceType) Byte() (byte, error) {
	switch et {
	case RelayEvidence:
		return 0, nil
	case ChallengeEvidence:
		return 1, nil
	default:
		return 0, fmt.Errorf("unrecognized EvidenceEncodable type")
	}
}

func EvidenceTypeFromString(evidenceType string) (et EvidenceType, err types.Error) {
	switch strings.ToLower(evidenceType) {
	case "relay":
		et = RelayEvidence
	case "challenge":
		et = ChallengeEvidence
	default:
		err = types.ErrInternal("type in the receipt query is not recognized: (relay or challenge)")
	}
	return
}
