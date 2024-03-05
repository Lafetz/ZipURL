package web

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ERRINVALIDPASSWORD = errors.New("incorrect password")
)

type Password struct {
	password string
}

func hashPassword(plaintextPassword string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
func matchPassword(plaintextPassword string, hash []byte) error {

	err := bcrypt.CompareHashAndPassword(hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return ERRINVALIDPASSWORD
		default:
			return err
		}
	}

	return err
}
