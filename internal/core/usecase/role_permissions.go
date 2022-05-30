package usecase

import (
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"strings"
)

type rolePermissionsService struct {
	rolePermissionsRepo ports.RolePermissionsRepository
	log                 ports.Logger
}

func NewRolePermissionsService(rolePermissionsRepo ports.RolePermissionsRepository, log ports.Logger) ports.RolePermissionsService {
	return &rolePermissionsService{
		rolePermissionsRepo: rolePermissionsRepo,
		log:                 log,
	}
}

func (s *rolePermissionsService) AddPermissionsToRole(rolePermissions *domain.RolePermissions) error {
	return s.rolePermissionsRepo.AddPermissionsToRole(rolePermissions)
}
func (s *rolePermissionsService) RemovePermissionFromRole(roleId string, permissionID string) error {
	return s.rolePermissionsRepo.RemovePermissionFromRole(strings.TrimSpace(roleId), strings.TrimSpace(permissionID))
}
func (s *rolePermissionsService) GetAllPermissionsForRole(roleId string) ([]domain.Permission, error) {
	return s.rolePermissionsRepo.GetAllPermissionsForRole(strings.TrimSpace(roleId))
}
