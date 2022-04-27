package ports

import "github.com/iBoBoTi/project_boiler_plate/internal/core/domain"

type RoleService interface {
	GetAllRoles() ([]domain.Role, error)
	GetRoleByID(id string) (*domain.Role, error)
	CreateRole(role *domain.Role) error
	AddPermission(id string, permission *domain.Permission) error
	RemovePermission(id string, permission *domain.Permission) error
	DeleteRole(id string) error
}
