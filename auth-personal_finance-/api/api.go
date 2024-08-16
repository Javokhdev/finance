package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"auth/api/handlers"
	"auth/api/middleware"
	_ "auth/api/docs"
)

// @title Authentication Service API
// @version 1.0
// @description API for Authentication Service
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Engine(handler *handlers.Handlers) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/register", handler.RegisterUser).Use(middleware.Middleware())
	router.POST("/login", handler.LoginUser).Use(middleware.Middleware())
	router.POST("/forgot-password", handler.ForgotPassword)
	router.POST("/reset-password", handler.ResetPassword)
	router.GET("/refresh-token", handler.RefreshToken)
	router.PUT("/change-role", handler.ChangeRole)

	user := router.Group("/user").Use(middleware.JWTMiddleware())
	{
		user.GET("/profiles", handler.GetProfile)
		user.PUT("/profiles", handler.EditProfile)
		user.PUT("/passwords", handler.ChangePassword)
		user.GET("/setting", handler.GetSetting)
		user.PUT("/setting", handler.EditSetting)
		user.DELETE("/", handler.DeleteUser)
	}

	return router
}
