package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
)

type rolePermissionsHandler struct {
	rolePermissionsService ports.RolePermissionsService
	log                    ports.Logger
}

func NewRolePermissionsHandler(rolePermissionsService ports.RolePermissionsService, log ports.Logger) ports.RolePermissionsHandler {
	return &rolePermissionsHandler{
		rolePermissionsService: rolePermissionsService,
		log:                    log,
	}
}

func (h *rolePermissionsHandler) AddPermissionsToRole(c *gin.Context) {
	rolePermissions := domain.RolePermissions{}
	if err := c.ShouldBindJSON(&rolePermissions); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.rolePermissionsService.AddPermissionsToRole(&rolePermissions); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (h *rolePermissionsHandler) RemovePermissionFromRole(c *gin.Context) {
	roleID := c.Param("role_id")
	permissionID := c.Param("permission_id")
	if err := h.rolePermissionsService.RemovePermissionFromRole(roleID, permissionID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (h *rolePermissionsHandler) GetAllPermissionsForRole(c *gin.Context) {
	roleID := c.Param("role_id")
	permissions, err := h.rolePermissionsService.GetAllPermissionsForRole(roleID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, permissions)

}
