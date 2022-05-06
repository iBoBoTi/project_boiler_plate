package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
)

// PermissionHandler is a permission handler interface for request and response handler .
type PermissionHandler interface {
	CreatePermission(c *gin.Context)
	DeletePermission(c *gin.Context)
	GetAllPermissions(c *gin.Context)
	GetPermissionByID(c *gin.Context)
}

// PermissionService is a service interface for the core to communicate with the adapters' permission handlers .
type PermissionService interface {
	CreatePermission(permission *domain.Permission) (*domain.Permission, error)
	DeletePermission(id string) error
	GetPermissionByID(id string) (*domain.Permission, error)
	GetAllPermissions() ([]domain.Permission, error)
}

// PermissionRepository is the interface for the core to communicate with the adapters' permission usecase.
type PermissionRepository interface {
	CreatePermission(permission *domain.Permission) (*domain.Permission, error)
	DeletePermission(id string) error
	GetPermissionByID(id string) (*domain.Permission, error)
	GetPermissionByTitle(title string) (*domain.Permission, error)
	GetAllPermissions() ([]domain.Permission, error)
}
