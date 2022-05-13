package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/adapters/api/response"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"net/http"
	"strconv"
)

type permissionHandler struct {
	permissionService ports.PermissionService
	logger            ports.Logger
}

func NewPermissionHandler(permissionService ports.PermissionService, logger ports.Logger) ports.PermissionHandler {
	return &permissionHandler{
		permissionService: permissionService,
		logger:            logger,
	}
}

func (h *permissionHandler) CreatePermission(c *gin.Context) {
	h.logger.Infof("Create Permission")
	permission := domain.Permission{}

	if err := c.ShouldBindJSON(&permission); err != nil {
		response.JSON(c, "invalid_request_body", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	resultPerm, err := h.permissionService.CreatePermission(&permission)
	if err != nil {
		response.JSON(c, "failed to create permission", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	response.JSON(c, "success in creating permission", http.StatusCreated, resultPerm, nil)
}

func (h *permissionHandler) GetPermissionByID(c *gin.Context) {
	h.logger.Infof("Get Permission By ID")
	permissionID := c.Param("id")
	if !domain.IsUUID(permissionID) {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
		return
	}

	permission, err := h.permissionService.GetPermissionByID(permissionID)
	if err != nil {
		h.logger.Errorf("Get Permission By ID: %s", err.Error())
		response.JSON(c, "invalid_input", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("Get Permission By ID: %s", permission.ID)
	response.JSON(c, "success in finding permission", http.StatusOK, permission, nil)
}

func (h *permissionHandler) GetAllPermissions(c *gin.Context) {
	h.logger.Infof("Get All Permissions")
	p := c.Query("page")
	if p == "" || p == "0" {
		p = "1"
	}
	page, err := strconv.Atoi(p)
	if err != nil {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
	}

	paginatedPermissions, err := h.permissionService.GetAllPermissions(page)
	if err != nil {
		h.logger.Errorf("Get All Permissions: %s", err.Error())
		response.JSON(c, "failed to find permissions", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("Get All Permissions")
	response.JSON(c, "success in finding permissions", http.StatusOK, paginatedPermissions, nil)
}

func (h *permissionHandler) DeletePermission(c *gin.Context) {
	h.logger.Infof("Delete Permission")
	id := c.Param("id")
	if !domain.IsUUID(id) {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
		return
	}

	if err := h.permissionService.DeletePermission(id); err != nil {
		h.logger.Errorf("Delete Permission: %s", err.Error())
		response.JSON(c, "failed to delete permission", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("Delete Permission with PermissionID: %s", id)
	response.JSON(c, "permission successfully deleted", http.StatusOK, nil, nil)
}
