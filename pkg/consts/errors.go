package consts

import "errors"

var (
	// ErrStartingCmp = errors.New(FmtCannotStart)
	ErrInvalidAffectedRows = errors.New("invalid number of affected rows")
	ErrInvalidInviteToken  = errors.New("invalid invite token")
	ErrInvalidAccessToken  = errors.New("invalid access token")
	ErrLoginAlreadyExists  = errors.New("login already exists")
	ErrInvalidCredentials  = errors.New("invalid credentials")
)
