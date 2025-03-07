package account

import "errors"

var (
	ErrInvalidIdFormat       = errors.New("invalid ID format")
	ErrAccountNotFound       = errors.New("account not found")
	ErrAccountCreate         = errors.New("something went wrong while creating the account")
	ErrAccountDeposit        = errors.New("something went wrong while depositing the amount")
	ErrAccountWithdraw       = errors.New("something went wrong while withdrawing the amount")
	ErrAccountAmountTransfer = errors.New("something went wrong while transfering the amount")
	ErrCloseAccount          = errors.New("something went wrong while closing the account")
)
