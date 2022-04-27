package api

import (
	"fmt"
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
		fmt.Println("i am here 2")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Role created successfully"})
}

func (h *roleHandler) GetRole(c *gin.Context) {}

func (h *roleHandler) GetRoles(c *gin.Context) {}

func (h *roleHandler) AddPermissionToRole(c *gin.Context) {}

func (h *roleHandler) RemovePermissionFromRole(c *gin.Context) {}

func (h *roleHandler) DeleteRole(c *gin.Context) {}
