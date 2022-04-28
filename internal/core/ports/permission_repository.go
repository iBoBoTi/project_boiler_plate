package ports

import "github.com/iBoBoTi/project_boiler_plate/internal/core/domain"

type PermissionRepository interface {
	CreatePermission(permission *domain.Permission) error
	DeletePermission(id string) error
	GetPermission(id string) (*domain.Permission, error)
	GetPermissionByTitle(title string) (*domain.Permission, error)
	GetAllPermissions() ([]domain.Permission, error)
}
