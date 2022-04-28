package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.roleService.CreateRole(&role); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Role created successfully"})
}

func (h *roleHandler) GetRole(c *gin.Context) {
	id := c.Param("id")
	role, err := h.roleService.GetRoleByID(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"role": role})
}

func (h *roleHandler) GetRoles(c *gin.Context) {
	roles, err := h.roleService.GetAllRoles()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"roles": roles})
}

func (h *roleHandler) AddPermissionToRole(c *gin.Context) {}

func (h *roleHandler) RemovePermissionFromRole(c *gin.Context) {}

func (h *roleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := h.roleService.DeleteRole(id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Role deleted successfully"})
}
