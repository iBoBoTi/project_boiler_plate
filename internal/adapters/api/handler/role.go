package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/adapters/api/response"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"net/http"
	"strconv"
)

type roleHandler struct {
	roleService ports.RoleService
	logger      ports.Logger
}

func NewRoleHandler(roleService ports.RoleService, logger ports.Logger) ports.RoleHandler {
	return &roleHandler{
		roleService: roleService,
		logger:      logger,
	}
}

func (h *roleHandler) CreateRole(c *gin.Context) {
	h.logger.Infof("creating role")

	role := domain.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		response.JSON(c, "invalid_request_body", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	resultRole, err := h.roleService.CreateRole(&role)
	if err != nil {
		h.logger.Errorf("creating role failed: %s", err.Error())
		response.JSON(c, "failed to create role", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("creating role with ID %s successful", resultRole.ID)
	response.JSON(c, "success creating role", http.StatusCreated, resultRole, nil)
}

func (h *roleHandler) GetRoleByID(c *gin.Context) {
	roleID := c.Param("id")
	h.logger.Infof("get role with id %s", roleID)

	if !domain.IsUUID(roleID) {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
		return
	}

	role, err := h.roleService.GetRoleByID(roleID)
	if err != nil {
		h.logger.Errorf("get role with id %s failed", roleID)
		response.JSON(c, "failed to find role", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("get role with id %s successful", role.ID)
	response.JSON(c, "success finding role", http.StatusOK, role, nil)
}

func (h *roleHandler) GetAllRoles(c *gin.Context) {
	p := c.Query("page")
	if p == "" || p == "0" {
		p = "1"
	}
	page, err := strconv.Atoi(p)
	if err != nil {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
	}

	h.logger.Infof("getting all roles")
	paginatedRoles, err := h.roleService.GetAllRoles(page)
	if err != nil {
		h.logger.Errorf("get all roles failed")
		response.JSON(c, "failed to find roles", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("get all roles successful")
	response.JSON(c, "success retrieving roles", http.StatusOK, paginatedRoles, nil)
}

func (h *roleHandler) DeleteRole(c *gin.Context) {
	roleID := c.Param("id")
	h.logger.Infof("delete role with id %s", roleID)

	if !domain.IsUUID(roleID) {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
		return
	}

	if err := h.roleService.DeleteRole(roleID); err != nil {
		h.logger.Errorf("delete role with id %s failed: %s", roleID, err.Error())
		response.JSON(c, "failed to delete role", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	h.logger.Infof("delete role with id %s successful", roleID)
	response.JSON(c, "success in deleting role", http.StatusOK, nil, nil)
}
