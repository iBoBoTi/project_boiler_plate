package usecase

import (
	"fmt"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/helpers"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"os"
	"strconv"
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

func (r *roleService) GetAllRoles(page int) (*helpers.Paginate, error) {
	limit, _ := strconv.Atoi(os.Getenv("PAGE_LIMIT"))

	paginate := helpers.NewPaginate(limit, page)
	paginate.Offset = (page - 1) * limit
	return r.roleRepo.GetAllRoles(paginate)
}
func (r *roleService) GetRoleByID(id string) (*domain.Role, error) {
	return r.roleRepo.GetRoleByID(strings.TrimSpace(id))
}

func (r *roleService) CreateRole(role *domain.Role) (*domain.Role, error) {
	role.ID = domain.NewUUID()
	role.Title = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(role.Title), " ", "_"))
	_, err := r.roleRepo.GetRoleByName(role.Title)
	if err != nil {
		return r.roleRepo.CreateRole(role)
	}
	return nil, fmt.Errorf("role title already exist")
}

func (r *roleService) DeleteRole(id string) error {
	return r.roleRepo.DeleteRole(strings.TrimSpace(id))
}
