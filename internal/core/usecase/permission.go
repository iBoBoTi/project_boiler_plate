package usecase

import (
	"fmt"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"strings"
)

type permissionService struct {
	permissionRepo ports.PermissionRepository
}

func NewPermissionService(permissionRepo ports.PermissionRepository) ports.PermissionService {
	return &permissionService{
		permissionRepo: permissionRepo,
	}
}

func (p *permissionService) CreatePermission(permission *domain.Permission) (*domain.Permission, error) {
	permission.ID = domain.NewUUID()
	permission.Title = strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(permission.Title), " ", "_"))
	_, err := p.permissionRepo.GetPermissionByTitle(permission.Title)
	if err != nil {
		return p.permissionRepo.CreatePermission(permission)
	}
	return nil, fmt.Errorf("permission title already exist")
}
func (p *permissionService) DeletePermission(id string) error {
	return p.permissionRepo.DeletePermission(strings.TrimSpace(id))
}
func (p *permissionService) GetPermissionByID(id string) (*domain.Permission, error) {
	return p.permissionRepo.GetPermissionByID(strings.TrimSpace(id))
}
func (p *permissionService) GetAllPermissions() ([]domain.Permission, error) {
	return p.permissionRepo.GetAllPermissions()
}
