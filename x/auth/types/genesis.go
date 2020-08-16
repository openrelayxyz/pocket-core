package types

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/pokt-network/pocket-core/codec/types"
)

// NewGenesisState - Create a new genesis state
func NewGenesisState(params Params, accounts Accounts) GenesisState {
	accs, err := PackAccounts(accounts)
	if err != nil {
		panic(err)
	}
	return GenesisState{
		Params:   params,
		Accounts: accs,
	}
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState(DefaultParams(), nil)
}

// ValidateGenesis performs basic validation of auth genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	accs, err := UnpackAccounts(data.Accounts)
	if err != nil {
		panic(err)
	}
	for _, account := range accs {
		if account.GetPubKey().PubKey() == nil {
			return fmt.Errorf("PubKey should never be nil")
		}
	}
	if data.Params.MaxMemoCharacters == 0 {
		return fmt.Errorf("invalid max memo characters: %d", data.Params.MaxMemoCharacters)
	}
	if data.Params.TxSigLimit == 0 {
		return fmt.Errorf("invalid tx signature limit: %d", data.Params.TxSigLimit)
	}
	if err := NewSupply(data.Supply).ValidateBasic(); err != nil {
		return err
	}
	return nil
}

// PackAccounts converts GenesisAccounts to Any slice
func PackAccounts(accounts Accounts) ([]*types.Any, error) {
	accountsAny := make([]*types.Any, len(accounts))
	for i, acc := range accounts {
		msg, ok := acc.(proto.Message)
		if !ok {
			return nil, fmt.Errorf("cannot proto marshal %T", acc)
		}
		any, err := types.NewAnyWithValue(msg)
		if err != nil {
			return nil, err
		}
		accountsAny[i] = any
	}

	return accountsAny, nil
}

// UnpackAccounts converts Any slice to GenesisAccounts
func UnpackAccounts(accountsAny []*types.Any) (Accounts, error) {
	accounts := make(Accounts, len(accountsAny))
	for i, any := range accountsAny {
		acc, ok := any.GetCachedValue().(BaseAccount)
		if !ok {
			return nil, fmt.Errorf("expected genesis account")
		}
		accounts[i] = &acc
	}

	return accounts, nil
}
