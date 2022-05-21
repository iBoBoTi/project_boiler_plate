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
	logger                 ports.Logger
}

func NewRolePermissionsHandler(rolePermissionsService ports.RolePermissionsService, logger ports.Logger) ports.RolePermissionsHandler {
	return &rolePermissionsHandler{
		rolePermissionsService: rolePermissionsService,
		logger:                 logger,
	}
}

func (h *rolePermissionsHandler) AddPermissionsToRole(c *gin.Context) {
	rolePermissions := domain.RolePermissions{}

	if err := c.ShouldBindJSON(&rolePermissions); err != nil {
		response.JSON(c, "invalid_request_body", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}
	h.logger.Infof("adding permissions to role with ID %s", rolePermissions.RoleID)

	if err := h.rolePermissionsService.AddPermissionsToRole(&rolePermissions); err != nil {
		h.logger.Errorf("adding permissions to role failed: %s", err.Error())
		response.JSON(c, "failed to add permissions", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("success adding permissions to role %s", rolePermissions.RoleID)
	response.JSON(c, "permissions added successfully", http.StatusOK, nil, nil)
}

func (h *rolePermissionsHandler) RemovePermissionFromRole(c *gin.Context) {
	roleID := c.Param("role_id")
	permissionID := c.Param("permission_id")

	h.logger.Infof("removing permission with ID %s from role with ID %s", permissionID, roleID)
	if !domain.IsUUID(roleID) && !domain.IsUUID(permissionID) {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
		return
	}
	if err := h.rolePermissionsService.RemovePermissionFromRole(roleID, permissionID); err != nil {
		h.logger.Errorf("removing permission with ID %s from role with ID %s failed: %s", permissionID, roleID, err.Error())
		response.JSON(c, "invalid_input", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("permission with ID %s removed successfully from role with ID %s", permissionID, roleID)
	response.JSON(c, "permission removed successfully", http.StatusOK, nil, nil)
}

func (h *rolePermissionsHandler) GetAllPermissionsForRole(c *gin.Context) {
	roleID := c.Param("role_id")

	if !domain.IsUUID(roleID) {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
		return
	}

	h.logger.Infof("get all permissions for role with ID %s", roleID)
	permissions, err := h.rolePermissionsService.GetAllPermissionsForRole(roleID)
	if err != nil {
		response.JSON(c, "invalid_input", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("getting all permissions for role with ID %s", roleID)
	response.JSON(c, "permissions gotten", http.StatusOK, permissions, nil)
}
