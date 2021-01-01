package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/mbcorecr module sentinel errors
var (
	ErrSample    = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrDateIssue = sdkerrors.Register(ModuleName, 1200, "Issue with a date")
)

// gov sentinel errors
var (
	ErrKeyring         = sdkerrors.Register(ModuleName, 1300, "Can't load Keyring")
	ErrKeyringAccount  = sdkerrors.Register(ModuleName, 1305, "Can't apply keyring method")
	ErrCryptConversion = sdkerrors.Register(ModuleName, 1310, "Con't convert key")
	ErrCipher          = sdkerrors.Register(ModuleName, 1320, "Can't load Keyring")
	ErrCreator         = sdkerrors.Register(ModuleName, 1400, "Transaction creator issue")
	ErrNewAccount      = sdkerrors.Register(ModuleName, 1500, "Can't create a new account")
)
