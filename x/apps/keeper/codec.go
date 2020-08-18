package keeper

import (
	"github.com/pokt-network/pocket-core/x/apps/types"
)

// EncodeApp - Encodes application to protobuff
func (k Keeper) EncodeApp(application types.Application) ([]byte, error) {
	return types.MarshalApplication(k.cdc, application)
}

// DecodeApp - Decodes application from protobuff
func (k Keeper) DecodeApp(ab []byte) (types.Application, error) {
	return types.UnmarshalApplication(k.cdc, ab)
}
