// Package usecase contains the business logic of the application - the services.
package usecase

import (
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
)

type userService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) GetUser() (*domain.User, error) {
	return nil, nil
}

func (u *userService) CreateUser() (*domain.User, error) {
	return nil, nil
}

func (u *userService) UpdateUser() (*domain.User, error) {
	return nil, nil
}

func (u *userService) DeleteUser() error {
	return nil
}
