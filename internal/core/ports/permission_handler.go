package ports

import "github.com/gin-gonic/gin"

type PermissionHandler interface {
	CreatePermission(c *gin.Context)
	DeletePermission(c *gin.Context)
	GetAllPermissions(c *gin.Context)
	GetPermissionByID(c *gin.Context)
}
