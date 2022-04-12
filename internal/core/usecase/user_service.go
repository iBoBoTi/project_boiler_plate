// Package usecase contains the business logic of the application - the services.
package usecase

import (
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
)

type service struct {
	userRepo ports.UserRepository
}

func NewService(userRepo ports.UserRepository) *service {
	return &service{
		userRepo: userRepo,
	}
}
