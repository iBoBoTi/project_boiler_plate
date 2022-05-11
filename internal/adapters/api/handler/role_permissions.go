package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/adapters/api/response"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"net/http"
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
		response.JSON(c, "invalid_request_body", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}
	if err := h.rolePermissionsService.AddPermissionsToRole(&rolePermissions); err != nil {
		response.JSON(c, "failed to add permissions", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "permissions added successfully", http.StatusOK, nil, nil)
}

func (h *rolePermissionsHandler) RemovePermissionFromRole(c *gin.Context) {
	roleID := c.Param("role_id")
	permissionID := c.Param("permission_id")
	if err := h.rolePermissionsService.RemovePermissionFromRole(roleID, permissionID); err != nil {
		response.JSON(c, "invalid_input", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "permission removed successfully", http.StatusOK, nil, nil)
}

func (h *rolePermissionsHandler) GetAllPermissionsForRole(c *gin.Context) {
	roleID := c.Param("role_id")
	permissions, err := h.rolePermissionsService.GetAllPermissionsForRole(roleID)
	if err != nil {
		response.JSON(c, "invalid_input", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "permissions gotten", http.StatusOK, permissions, nil)

}
