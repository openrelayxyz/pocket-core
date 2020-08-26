package keeper

import (
	"fmt"

	"github.com/pokt-network/pocket-core/x/auth/exported"
	"github.com/pokt-network/pocket-core/x/auth/types"

	sdk "github.com/pokt-network/pocket-core/types"
)

// SendCoinsFromModuleToAccount transfers coins from a ModuleAccount to an Address
func (k Keeper) SendCoinsFromModuleToAccount(ctx sdk.Ctx, senderModule string,
	recipientAddr sdk.Address, amt sdk.Coins) sdk.Error {

	sender := k.GetModuleAccount(ctx, senderModule)
	if sender == nil || sender.GetAddress() == nil {
		return sdk.ErrUnknownAddress(fmt.Sprintf("module account %s does not exist", senderModule))
	}
	recipient := k.GetAccount(ctx, recipientAddr)
	if recipient == nil {
		var err error
		recipient, err = k.NewAccountWithAddress(ctx, recipientAddr)
		if err != nil {
			return sdk.ErrInternal(err.Error())
		}
	}

	return k.sendCoins(ctx, sender, recipient, amt)
}

// SendCoinsFromModuleToModule transfers coins from a ModuleAccount to another
func (k Keeper) SendCoinsFromModuleToModule(ctx sdk.Ctx, senderModule, recipientModule string, amt sdk.Coins) sdk.Error {
	sender := k.GetModuleAccount(ctx, senderModule)
	if sender == nil || sender.GetAddress() == nil {
		return sdk.ErrUnknownAddress(fmt.Sprintf("module account %s does not exist", senderModule))
	}

	recipient := k.GetModuleAccount(ctx, recipientModule)
	// create the account if it doesn't yet exist
	if recipient == nil {
		return sdk.ErrModuleAccountCreate(fmt.Sprintf("module account %s isn't able to be created", recipientModule))
	}

	return k.sendCoins(ctx, sender, recipient, amt)
}

// SendCoinsFromAccountToModule transfers coins from an Address to a ModuleAccount
func (k Keeper) SendCoinsFromAccountToModule(ctx sdk.Ctx, senderAddr sdk.Address,
	recipientModule string, amt sdk.Coins) sdk.Error {
	sender := k.GetAccount(ctx, senderAddr)
	if sender == nil || sender.GetAddress() == nil {
		return sdk.ErrUnknownAddress(fmt.Sprintf("account %s does not exist", senderAddr))
	}

	// create the account if it doesn't yet exist
	recipient := k.GetModuleAccount(ctx, recipientModule)
	// create the account if it doesn't yet exist
	if recipient == nil {
		return sdk.ErrModuleAccountCreate(fmt.Sprintf("module account %s isn't able to be created", recipientModule))
	}

	return k.sendCoins(ctx, sender, recipient, amt)
}

// MintCoins creates new coins from thin air and adds it to the module account.
// Panics if the name maps to a non-minter module account or if the amount is invalid.
func (k Keeper) MintCoins(ctx sdk.Ctx, moduleName string, amt sdk.Coins) sdk.Error {

	// create the account if it doesn't yet exist
	acc := k.GetModuleAccount(ctx, moduleName)
	if acc == nil {
		return sdk.ErrUnknownAddress(fmt.Sprintf("module account %s does not exist", moduleName))
	}

	if !acc.HasPermission(types.Minter) {
		return sdk.ErrForbidden(fmt.Sprintf("module account %s does not have permissions to mint tokens", moduleName))
	}

	// Actually Add coins to acc
	_, err := k.AddCoins(ctx, acc, amt)
	if err != nil {
		return sdk.ErrInternal(err.Error())
	}

	// update total supply
	supply := k.GetSupply(ctx)
	supply = supply.Inflate(amt)

	k.SetSupply(ctx, supply)

	logger := k.Logger(ctx)
	logger.Debug(fmt.Sprintf("minted %s from %s module account", amt.String(), moduleName))

	return nil
}

// BurnCoins burns coins deletes coins from the balance of the module account.
// Panics if the name maps to a non-burner module account or if the amount is invalid.
func (k Keeper) BurnCoins(ctx sdk.Ctx, moduleName string, amt sdk.Coins) sdk.Error {

	// create the account if it doesn't yet exist
	acc := k.GetModuleAccount(ctx, moduleName)
	if acc == nil {
		return sdk.ErrUnknownAddress(fmt.Sprintf("module account %s does not exist", moduleName))
	}

	if !acc.HasPermission(types.Burner) {
		return sdk.ErrModuleAccountCreate(fmt.Sprintf("module account %s does not have permissions to burn tokens", moduleName))
	}

	// TODO actually subtract coins from Acc
	_, err := k.SubtractCoins(ctx, acc, amt)
	if err != nil {
		return sdk.ErrInternal(err.Error())
	}

	// update total supply
	supply := k.GetSupply(ctx)
	supply = supply.Deflate(amt)
	k.SetSupply(ctx, supply)

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("burned %s from %s module account", amt.String(), moduleName))

	return nil
}

// SendCoins moves coins from one account to another
func (k Keeper) SendCoins(ctx sdk.Ctx, fromAddr sdk.Address, toAddress sdk.Address, amt sdk.Coins) sdk.Error {
	sender := k.GetAccount(ctx, fromAddr)
	if sender == nil || sender.GetAddress() == nil {
		return sdk.ErrUnknownAddress(fmt.Sprintf("account %s does not exist", fromAddr.String()))
	}
	recipient := k.GetAcc(ctx, toAddress)
	t := recipient == nil
	if t {
		var err error
		recipient, err = k.NewAccountWithAddress(ctx, toAddress)
		if err != nil {
			return sdk.ErrInternal(err.Error())
		}
	}

	return k.sendCoins(ctx, sender, recipient, amt)
}

// SendCoins moves coins from one account to another
func (k Keeper) sendCoins(ctx sdk.Ctx, from exported.Account, to exported.Account, amt sdk.Coins) sdk.Error {
	_, err := k.SubtractCoins(ctx, from, amt)
	if err != nil {
		return err
	}
	_, err = k.AddCoins(ctx, to, amt)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransfer,
			sdk.NewAttribute(types.AttributeKeyRecipient, to.GetAddress().String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amt.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.AttributeKeySender, from.GetAddress().String()),
		),
	})

	return nil
}

// SubtractCoins subtracts amt from the coins at the addr.
func (k Keeper) SubtractCoins(ctx sdk.Ctx, acc exported.Account, amt sdk.Coins) (sdk.Coins, sdk.Error) {
	if !amt.IsValid() {
		return nil, sdk.ErrInvalidCoins(amt.String())
	}

	oldCoins := acc.GetCoins()
	spendableCoins := acc.SpendableCoins(ctx.BlockHeader().Time)

	// For non-vesting accounts, spendable coins will simply be the original coins.
	// So the check here is sufficient instead of subtracting from oldCoins.
	_, hasNeg := spendableCoins.SafeSub(amt)
	if hasNeg {
		return amt, sdk.ErrInsufficientCoins(
			fmt.Sprintf("insufficient account funds; %s < %s", spendableCoins, amt),
		)
	}

	newCoins := oldCoins.Sub(amt) // should not panic as spendable coins was already checked

	acc.SetCoins(newCoins)
	k.SetAccount(ctx, acc)

	return newCoins, nil
}

// AddCoins adds amt to the coins at the addr.
func (k Keeper) AddCoins(ctx sdk.Ctx, acc exported.Account, amt sdk.Coins) (sdk.Coins, sdk.Error) {
	if !amt.IsValid() {
		return nil, sdk.ErrInvalidCoins(amt.String())
	}

	oldCoins := acc.GetCoins()
	newCoins := oldCoins.Add(amt)

	if newCoins.IsAnyNegative() {
		return amt, sdk.ErrInsufficientCoins(
			fmt.Sprintf("insufficient account funds; %s < %s", oldCoins, amt),
		)
	}
	acc.SetCoins(newCoins)
	k.SetAccount(ctx, acc)

	return newCoins, nil
}

// SetCoins sets the coins at the addr.
func (k Keeper) SetCoins(ctx sdk.Ctx, addr sdk.Address, amt sdk.Coins) sdk.Error {
	if !amt.IsValid() {
		return sdk.ErrInvalidCoins(amt.String())
	}

	acc := k.GetAcc(ctx, addr)
	if acc == nil {
		var err error
		acc, err = k.NewAccountWithAddress(ctx, addr)
		if err != nil {
			return sdk.ErrInternal(err.Error())
		}
	}
	err := acc.SetCoins(amt)
	if err != nil {
		return sdk.ErrInternal(err.Error())
	}

	k.SetAccount(ctx, acc)
	return nil
}

// GetCoins returns the coins at the addr.
func (k Keeper) GetCoins(ctx sdk.Ctx, addr sdk.Address) sdk.Coins {
	acc := k.GetAcc(ctx, addr)
	if acc == nil {
		return sdk.NewCoins()
	}
	return acc.GetCoins()
}

// HasCoins returns whether or not an account has at least amt coins.
func (k Keeper) HasCoins(ctx sdk.Ctx, addr sdk.Address, amt sdk.Coins) bool {
	return k.GetCoins(ctx, addr).IsAllGTE(amt)
}

func (k Keeper) GetFee(ctx sdk.Ctx, msg sdk.Msg) sdk.Int {
	return k.GetParams(ctx).FeeMultiplier.GetFee(msg)
}
