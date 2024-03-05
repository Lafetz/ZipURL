package web

import (
	"github.com/google/uuid"
)

type userToken struct {
	Id       uuid.UUID
	Email    string
	Username string
}
