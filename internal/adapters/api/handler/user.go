package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
)

type userHandler struct {
	userService ports.UserService
	logger      ports.Logger
}

func NewUserHandler(userService ports.UserService, l ports.Logger) ports.UserHandler {
	return &userHandler{
		userService: userService,
		logger:      l,
	}
}

func (h *userHandler) GetUserByID(c *gin.Context) {}

func (h *userHandler) CreateUser(c *gin.Context) {}

func (h *userHandler) UpdateUser(c *gin.Context) {}

func (h *userHandler) DeleteUser(c *gin.Context) {}

func (h *userHandler) GetUsersList(c *gin.Context) {}
