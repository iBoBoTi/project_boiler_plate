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
	h.logger.Infof("creating permission")
	var permission domain.Permission

	if err := c.ShouldBindJSON(&permission); err != nil {
		response.JSON(c, "invalid_request_body", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	resultPerm, err := h.permissionService.CreatePermission(&permission)
	if err != nil {
		h.logger.Errorf("Creating Permission Failed: %s", err.Error())
		response.JSON(c, "failed to create permission", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("Creating Permission With ID %s Successful", resultPerm.ID)
	response.JSON(c, "success in creating permission", http.StatusCreated, resultPerm, nil)
}

func (h *permissionHandler) GetPermissionByID(c *gin.Context) {
	permissionID := c.Param("id")
	if !domain.IsUUID(permissionID) {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
		return
	}

	h.logger.Infof("get permission by ID %v", permissionID)
	permission, err := h.permissionService.GetPermissionByID(permissionID)
	if err != nil {
		h.logger.Errorf("get permission by ID %s failed: %s", permissionID, err.Error())
		response.JSON(c, "permission not found", http.StatusNotFound, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("get permission by ID: %s successful", permission.ID)
	response.JSON(c, "success in finding permission", http.StatusOK, permission, nil)
}

func (h *permissionHandler) GetAllPermissions(c *gin.Context) {
	h.logger.Infof("get all permissions")
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
		h.logger.Errorf("get all permissions failed: %s", err.Error())
		response.JSON(c, "failed to find permissions", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("get all permissions page %v of %v", paginatedPermissions.Page, paginatedPermissions.TotalPages)
	response.JSON(c, "success in finding permissions", http.StatusOK, paginatedPermissions, nil)
}

func (h *permissionHandler) DeletePermission(c *gin.Context) {
	permissionID := c.Param("id")

	h.logger.Infof("delete permission with ID %s", permissionID)

	if !domain.IsUUID(permissionID) {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
		return
	}

	if err := h.permissionService.DeletePermission(permissionID); err != nil {
		h.logger.Errorf("delete permission with ID %s failed: %s", permissionID, err.Error())
		response.JSON(c, "failed to delete permission, permission not found", http.StatusNotFound, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("delete permission with permissionID: %s successful", permissionID)
	response.JSON(c, "permission successfully deleted", http.StatusOK, nil, nil)
}
