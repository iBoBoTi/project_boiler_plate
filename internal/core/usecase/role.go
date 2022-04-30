package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"strings"
)

type roleService struct {
	roleRepo ports.RoleRepository
}

func NewRoleService(roleRepo ports.RoleRepository) ports.RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}

func (r *roleService) GetAllRoles() ([]domain.Role, error) {
	return r.roleRepo.GetAllRoles()
}
func (r *roleService) GetRoleByID(id string) (*domain.Role, error) {
	return r.roleRepo.GetRoleByID(strings.TrimSpace(id))
}

func (r *roleService) CreateRole(role *domain.Role) error {
	role.ID = uuid.New().String()
	role.Title = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(role.Title), " ", "_"))
	_, err := r.roleRepo.GetRoleByName(role.Title)
	if err != nil {
		return r.roleRepo.CreateRole(role)
	}
	return fmt.Errorf("role title already exist")
}

func (r *roleService) DeleteRole(id string) error {
	return r.roleRepo.DeleteRole(strings.TrimSpace(id))
}
