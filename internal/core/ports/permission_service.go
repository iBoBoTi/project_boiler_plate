package ports

import "github.com/iBoBoTi/project_boiler_plate/internal/core/domain"

type PermissionService interface {
	CreatePermission(permission *domain.Permission) error
	DeletePermission(id string) error
	GetPermissionByID(id string) (*domain.Permission, error)
	GetAllPermissions() ([]domain.Permission, error)
}
