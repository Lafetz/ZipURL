package services

import (
	"errors"
)

var (
	ErrUserNotFound   = errors.New("user not Found")
	ErrUsernameUnique = errors.New("an accoutn with this Username exists")
	ErrDelete         = errors.New("failed to Delete User")
	ErrEmailUnique    = errors.New("an account with this Email exists")
)
var (
	ErrMaxUrl = errors.New("max amount of url reached")
)

//+ strconv.Itoa(maxurl)
