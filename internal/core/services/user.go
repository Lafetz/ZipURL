package services

import (
	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/lafetz/url-shortner/internal/core/ports"
)

//	type UserService interface {
//		GetUser(string) (*domain.User, error)
//		AddUser(*domain.User) (*domain.User, error)
//		UpdateUser(*domain.User) error
//		DeleteUser(uuid.UUID) error
//	}
type UserService struct {
	repo ports.UserRepository
}

func (srv *UserService) GetUser(username string) (*domain.User, error) {
	return srv.repo.GetUser(username)
}
func (srv *UserService) AddUser(user *domain.User) (*domain.User, error) {

	return srv.repo.AddUser(user)
}
func (srv *UserService) UpdateUser(user *domain.User) error {

	return srv.repo.UpdateUser(user)
}
func (srv *UserService) DeleteUser(id uuid.UUID) error {
	return srv.repo.DeleteUser(id)
}
