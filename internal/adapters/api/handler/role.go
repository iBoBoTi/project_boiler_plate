package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/adapters/api/response"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"net/http"
)

type roleHandler struct {
	roleService       ports.RoleService
	permissionService ports.PermissionService
	logger            ports.Logger
}

func NewRoleHandler(roleService ports.RoleService, permissionService ports.PermissionService, logger ports.Logger) ports.RoleHandler {
	return &roleHandler{
		roleService:       roleService,
		permissionService: permissionService,
		logger:            logger,
	}
}

func (h *roleHandler) CreateRole(c *gin.Context) {
	role := domain.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		response.JSON(c, "invalid_request_body", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}
	if err := h.roleService.CreateRole(&role); err != nil {
		response.JSON(c, "invalid_input", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "role created successfully", http.StatusCreated, nil, nil)
}

func (h *roleHandler) GetRole(c *gin.Context) {
	id := c.Param("id")
	role, err := h.roleService.GetRoleByID(id)
	if err != nil {
		response.JSON(c, "invalid_input", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "role gotten", http.StatusCreated, role, nil)
}

func (h *roleHandler) GetRoles(c *gin.Context) {
	roles, err := h.roleService.GetAllRoles()
	if err != nil {
		response.JSON(c, "invalid_input", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "roles gotten", http.StatusCreated, roles, nil)
}

func (h *roleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := h.roleService.DeleteRole(id); err != nil {
		response.JSON(c, "invalid_input", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "roles deleted successfully", http.StatusCreated, nil, nil)
}
