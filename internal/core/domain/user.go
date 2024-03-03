package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Username  string
	Email     string
	Password  []byte
	CreatedAt time.Time
}
