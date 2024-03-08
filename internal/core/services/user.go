package services

import (
	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/lafetz/url-shortner/internal/core/ports"
)

type UserServiceApi interface {
	GetUser(string) (*domain.User, error)
	AddUser(*domain.User) (*domain.User, error)
	DeleteUser(uuid.UUID) error
}
type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (srv *UserService) GetUser(username string) (*domain.User, error) {
	return srv.repo.GetUser(username)
}
func (srv *UserService) AddUser(user *domain.User) (*domain.User, error) {
	//
	return srv.repo.AddUser(user)
}

func (srv *UserService) DeleteUser(id uuid.UUID) error {
	return srv.repo.DeleteUser(id)
}
