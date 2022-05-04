package server

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/project_boiler_plate/internal/adapters/api"
	"github.com/iBoBoTi/project_boiler_plate/internal/adapters/repositories/psql"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/usecase"
	"github.com/iBoBoTi/project_boiler_plate/pkg/database"
	"github.com/iBoBoTi/project_boiler_plate/pkg/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ginServer struct {
	log    ports.Logger
	router *router.Router
}

func newGinServer(l ports.Logger, r *router.Router) *ginServer {
	return &ginServer{
		log:    l,
		router: r,
	}
}

func (s *ginServer) setAppHandlers(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	db, _ := database.NewDatabaseFactory(database.InstancePostgres)

	//Permission
	permissionRepo := psql.NewPermissionRepository(db.Pool)
	permissionService := usecase.NewPermissionService(permissionRepo)
	permissionHandler := api.NewPermissionHandler(permissionService)

	permissionRouter := v1.Group("/permissions")
	permissionRouter.GET("/:id", permissionHandler.GetPermissionByID)
	permissionRouter.POST("/", permissionHandler.CreatePermission)
	permissionRouter.GET("/", permissionHandler.GetAllPermissions)
	permissionRouter.DELETE("/:id", permissionHandler.DeletePermission)

	// Role
	roleRepo := psql.NewRoleRepository(db.Pool)
	roleService := usecase.NewRoleService(roleRepo)
	roleHandler := api.NewRoleHandler(roleService, permissionService, s.log)

	roleRouter := v1.Group("/roles")
	roleRouter.GET("/:id", roleHandler.GetRole)
	roleRouter.POST("/", roleHandler.CreateRole)
	roleRouter.GET("/", roleHandler.GetRoles)
	roleRouter.DELETE("/:id", roleHandler.DeleteRole)

	// Role Permissions
	rolePermissionRepo := psql.NewRolePermissionsRepository(db.Pool, s.log)
	rolePermissionService := usecase.NewRolePermissionsService(rolePermissionRepo, s.log)
	rolePermissionHandler := api.NewRolePermissionsHandler(rolePermissionService, s.log)

	rolePermissionRouter := v1.Group("/role-permissions")
	rolePermissionRouter.POST("/", rolePermissionHandler.AddPermissionsToRole)
	rolePermissionRouter.GET("/:role_id", rolePermissionHandler.GetAllPermissionsForRole)
	rolePermissionRouter.DELETE("/:role_id/:permission_id", rolePermissionHandler.RemovePermissionFromRole)

	//User
	userRepo := psql.NewUserRepository(db.Pool)
	userService := usecase.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService, s.log)

	userRouter := v1.Group("/users")
	userRouter.GET("/:id", userHandler.GetUserByID)
	userRouter.POST("/", userHandler.CreateUser)
	userRouter.GET("/", userHandler.GetUsersList)
	userRouter.PUT("/:id", userHandler.UpdateUser)
	userRouter.DELETE("/:id", userHandler.DeleteUser)

}

func (s *ginServer) setupRouter() *gin.Engine {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "test" {
		r := gin.New()
		s.setAppHandlers(r)
		return r
	}

	r := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	// setup cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
	s.setAppHandlers(r)
	return r
}

func (s *ginServer) Run() {
	gin.SetMode(gin.ReleaseMode)
	gin.Recovery()

	r := s.setupRouter()
	port := os.Getenv("SERVER_PORT")

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%v", port),
		Handler:      r,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s.log.WithFields(ports.Fields{"port": port}).Infof("Starting HTTP Server")
		if err := server.ListenAndServe(); err != nil {
			s.log.WithError(err).Fatalln("Error starting HTTP server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		s.log.WithError(err).Fatalln("Server Shutdown Failed")
	}

	s.log.Infof("Service down")
}
