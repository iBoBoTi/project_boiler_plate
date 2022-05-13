package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/helpers"
)

// RoleHandler is a role handler interface for request and response handler .
type RoleHandler interface {
	GetRoleByID(c *gin.Context)
	GetAllRoles(c *gin.Context)
	CreateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}

// RoleService is a service interface for the core to communicate with the adapters' role handlers .
type RoleService interface {
	GetAllRoles(page int) (*helpers.Paginate, error)
	GetRoleByID(id string) (*domain.Role, error)
	CreateRole(role *domain.Role) (*domain.Role, error)
	DeleteRole(id string) error
}

// RoleRepository is the interface for the core to communicate with the adapters' role usecase.
type RoleRepository interface {
	GetRoleByName(name string) (*domain.Role, error)
	GetAllRoles(paginate *helpers.Paginate) (*helpers.Paginate, error)
	GetRoleByID(id string) (*domain.Role, error)
	CreateRole(role *domain.Role) (*domain.Role, error)
	DeleteRole(id string) error
}
