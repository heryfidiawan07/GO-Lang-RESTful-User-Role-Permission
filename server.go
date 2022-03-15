package main

import (
	"app/config"
	"app/controller"
	"app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	config.InitDB()
	// defer config.DB.Close()
	gotenv.Load()

	router := gin.Default()

	v1 := router.Group("/api/v1/")
	{
		// Home
		v1.GET("/", controller.Home)

		// Auth
		v1.GET("/auth/:provider", controller.RedirectHandler)
		v1.GET("/auth/:provider/callback", controller.CallbackHandler)
		//
		v1.POST("/auth/login", controller.Login)
		v1.POST("/auth/register", controller.Register)
		v1.POST("/auth/refresh-token", controller.RefreshToken)

		// Me
		v1.GET("/me", middleware.Auth("me"), controller.Me)
		v1.PUT("/changepassword", middleware.Auth("me"), controller.ChangePassword)
		// Revoke Token
		v1.PUT("/revoke", middleware.Auth("me"), controller.RevokeRefreshToken)

		// Concurency
		v1.GET("/concurency", middleware.Auth("me"), controller.ConcurentIndex)

		// File
		// Set a lower memory limit for multipart forms (default is 32 MiB)
		router.MaxMultipartMemory = 8 << 20 // 8 MiB
		v1.POST("/upload/:disk", controller.Upload)
		v1.GET("/storage/:filename", controller.FileStream)
		v1.GET("/encode/:filename", controller.Encode)

		// User
		user := v1.Group("/user")
		{
			user.GET("/", middleware.Auth("user-index"), controller.UserIndex)
			user.POST("/", middleware.Auth("user-store"), controller.UserStore)
			user.PUT("/:id", middleware.Auth("user-update"), controller.UserUpdate)
			user.GET("/:id", middleware.Auth("user-show"), controller.UserShow)
			user.DELETE("/:id", middleware.Auth("user-delete"), controller.UserDelete)
		}

		// Role
		role := v1.Group("/role")
		{
			// Json POST
			role.GET("/", middleware.Auth("role-index"), controller.RoleIndex)
			role.POST("/", middleware.Auth("role-store"), controller.RoleStore)
			role.PUT("/:id", middleware.Auth("role-update"), controller.RoleUpdate)
			role.GET("/:id", middleware.Auth("role-show"), controller.RoleShow)
			role.DELETE("/:id", middleware.Auth("role-delete"), controller.RoleDelete)
		}

		// Permission
		permission := v1.Group("/permission")
		{
			permission.GET("/", middleware.Auth("except"), controller.PermissionIndex)
			permission.POST("/", middleware.Auth("except"), controller.PermissionStore)
		}
	}

	router.Run("127.0.0.1:8000")
}
