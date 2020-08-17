package keeper

import "crypto"

// EncodeApp - Encodes application to protobuff
func (k Keeper) EncodeApp(application Application) (bytes, error) {
	app := ApplicationEncodable{
		Address:                 application.Address,
		PublicKey:               application.PublicKey.String(),
		Jailed:                  application.Jailed,
		Status:                  application.Status,
		Chains:                  application.Chains,
		StakedTokens:            application.StakedTokens,
		MaxRelays:               application.MaxRelays,
		UnstakingCompletionTime: application.UnstakingCompletionTime,
	}
	return k.cdc.MarshalBinaryLengthPrefixed(app)
}

// DecodeApp - Decodes application from protobuff
func (k Keeper) DecodeApp(ba []byte) (Application, error) {
	application := ApplicationEncodable{}
	err = k.cdc.UnmarshalBinaryLengthPrefixed(appBytes, &application)
	app := Application{
		Address:                 application.Address,
		PublicKey:               crypto.NewPublicKey(application.PublicKey),
		Jailed:                  application.Jailed,
		Status:                  application.Status,
		Chains:                  application.Chains,
		StakedTokens:            application.StakedTokens,
		MaxRelays:               application.MaxRelays,
		UnstakingCompletionTime: application.UnstakingCompletionTime,
	}
	return k.cdc.MarshalBinaryLengthPrefixed(app)
}
