package ports

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetUserByID(c *gin.Context)
	GetUsersList(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
