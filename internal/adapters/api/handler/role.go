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
		response.JSON(c, "failed to create role", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "success creating role", http.StatusCreated, nil, nil)
}

func (h *roleHandler) GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	role, err := h.roleService.GetRoleByID(id)
	if err != nil {
		response.JSON(c, "failed to find role", http.StatusNotFound, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "success finding role", http.StatusOK, role, nil)
}

func (h *roleHandler) GetAllRoles(c *gin.Context) {
	roles, err := h.roleService.GetAllRoles()
	if err != nil {
		response.JSON(c, "invalid_input", http.StatusNotFound, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "roles gotten", http.StatusOK, roles, nil)
}

func (h *roleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := h.roleService.DeleteRole(id); err != nil {
		response.JSON(c, "failed to delete role", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "success in deleting role", http.StatusOK, nil, nil)
}
