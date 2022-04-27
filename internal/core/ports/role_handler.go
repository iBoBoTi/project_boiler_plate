package ports

import (
	"github.com/gin-gonic/gin"
)

type RoleHandler interface {
	GetRole(c *gin.Context)
	GetRoles(c *gin.Context)
	CreateRole(c *gin.Context)
	AddPermissionToRole(c *gin.Context)
	RemovePermissionFromRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}
