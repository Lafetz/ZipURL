package domain

import (
	"time"

	"github.com/google/uuid"
)

type Url struct {
	Id             uuid.UUID
	UserId         uuid.UUID
	ShortUrl       string
	OriginalUrl    string
	ExpirationDate time.Time
}
