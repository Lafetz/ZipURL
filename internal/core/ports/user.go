package ports

import (
	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
)

type UserRepository interface {
	GetUser(username string) (*domain.User, error)
	AddUser(*domain.User) (*domain.User, error)
	DeleteUser(uuid.UUID) error
}
