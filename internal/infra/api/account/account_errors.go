package account

import "errors"

var (
	ErrInvalidIdFormat = errors.New("invalid ID format")
	ErrAccountNotFound = errors.New("account not found")
	ErrAccountCreate   = errors.New("something went wrong while creating the account")
	ErrAccountDeposit  = errors.New("something went wrong while depositing the amount")
)
