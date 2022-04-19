package ports

import "github.com/iBoBoTi/project_boiler_plate/internal/core/domain"

// UserService is a service interface for the core to communicate with the adapters handlers .
type UserService interface {
	GetUser() (*domain.User, error)
	CreateUser() (*domain.User, error)
	UpdateUser() (*domain.User, error)
	DeleteUser() error
}
