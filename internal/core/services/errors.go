package services

import "errors"

var (
	ErrUserNotFound   = errors.New("user not Found")
	ErrUsernameUnique = errors.New("an accoutn with this Username exists")
	// ErrUpdateUser     = errors.New("failed to Update User")
	ErrDelete      = errors.New("failed to Delete User")
	ErrEmailUnique = errors.New("an account with this Email exists")
	// ErrUserVerfied    = errors.New("email not Verified")
)
