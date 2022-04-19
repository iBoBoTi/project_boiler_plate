// Package ports contains interfaces for the adapters to communicate to the core.
package ports

import "github.com/iBoBoTi/project_boiler_plate/internal/core/domain"

// UserRepository is the interface for the core to communicate with the adapters usecase.
type UserRepository interface {
	GetUser() (*domain.User, error)
	CreateUser() (*domain.User, error)
	UpdateUser() (*domain.User, error)
	DeleteUser() error
}
