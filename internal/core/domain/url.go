package domain

import (
	"time"

	"github.com/google/uuid"
)

type Url struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	ShortUrl    string
	OriginalUrl string
	CreatedAt   time.Time
}

func NewUrl(userId uuid.UUID, originalUrl string) *Url {
	return &Url{
		Id:          uuid.New(),
		UserId:      userId,
		OriginalUrl: originalUrl,
	}
}
