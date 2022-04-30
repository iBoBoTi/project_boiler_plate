package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
)

// UserHandler is a user handler interface for request and response handler .
type UserHandler interface {
	GetUserByID(c *gin.Context)
	GetUsersList(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

// UserService is a service interface for the core to communicate with the adapters' user handlers .
type UserService interface {
	GetUser() (*domain.User, error)
	CreateUser() (*domain.User, error)
	UpdateUser() (*domain.User, error)
	DeleteUser() error
}

// UserRepository is the interface for the core to communicate with the adapters' user usecase.
type UserRepository interface {
	GetUser() (*domain.User, error)
	CreateUser() (*domain.User, error)
	UpdateUser() (*domain.User, error)
	DeleteUser() error
}
