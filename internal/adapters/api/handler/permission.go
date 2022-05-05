package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/adapters/api/response"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"net/http"
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
		response.JSON(c, "invalid_request_body", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}
	if err := h.permissionService.CreatePermission(&permission); err != nil {
		response.JSON(c, "invalid_input", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "permission successfully created", http.StatusCreated, nil, nil)
}

func (h *permissionHandler) GetPermissionByID(c *gin.Context) {
	permissionID := c.Param("id")
	permission, err := h.permissionService.GetPermissionByID(permissionID)
	if err != nil {
		response.JSON(c, "invalid_input", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "permission gotten", http.StatusOK, permission, nil)
}

func (h *permissionHandler) GetAllPermissions(c *gin.Context) {
	permissions, err := h.permissionService.GetAllPermissions()
	if err != nil {
		response.JSON(c, "invalid_input", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "permissions gotten", http.StatusOK, permissions, nil)
}

func (h *permissionHandler) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	if err := h.permissionService.DeletePermission(id); err != nil {
		response.JSON(c, "invalid_input", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "permission successfully deleted", http.StatusOK, nil, nil)
}
