package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
)

type permissionHandler struct {
	permissionService ports.PermissionService
}

func NewPermissionHandler(permissionService ports.PermissionService) ports.PermissionHandler {
	return &permissionHandler{
		permissionService: permissionService,
	}
}

func (h *permissionHandler) CreatePermission(c *gin.Context) {
	permission := domain.Permission{}
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.permissionService.CreatePermission(&permission); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "permission successfully created"})
}

func (h *permissionHandler) GetPermissionByID(c *gin.Context) {
	permissionID := c.Param("id")
	permission, err := h.permissionService.GetPermissionByID(permissionID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"permission": permission})
}

func (h *permissionHandler) GetAllPermissions(c *gin.Context) {
	permissions, err := h.permissionService.GetAllPermissions()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"permissions": permissions})
}

func (h *permissionHandler) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	if err := h.permissionService.DeletePermission(id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "permission successfully deleted"})
}
