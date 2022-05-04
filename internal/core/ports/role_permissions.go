package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
)

type RolePermissionsHandler interface {
	AddPermissionsToRole(c *gin.Context)
	RemovePermissionFromRole(c *gin.Context)
	GetAllPermissionsForRole(c *gin.Context)
}

type RolePermissionsService interface {
	AddPermissionsToRole(rolePermissions *domain.RolePermissions) error
	RemovePermissionFromRole(roleId string, permissionID string) error
	GetAllPermissionsForRole(roleId string) ([]domain.Permission, error)
}

type RolePermissionsRepository interface {
	AddPermissionsToRole(rolePermissions *domain.RolePermissions) error
	RemovePermissionFromRole(roleId string, permissionID string) error
	GetAllPermissionsForRole(roleId string) ([]domain.Permission, error)
}
