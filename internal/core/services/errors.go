package services

import (
	"errors"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrUsernameUnique = errors.New("an account with this username exists")
	ErrDelete         = errors.New("failed to Delete user")
	ErrEmailUnique    = errors.New("an account with this email exists")
)
var (
	ErrUrlNotFound        = errors.New("url not found")
	ErrUrlDepulicateRetry = errors.New("retry")
	ErrDepulicateShortUrl = errors.New("short url depulicate")
)
