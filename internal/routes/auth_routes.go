package routes

import (
	"github.com/Eduardo-Brehm/mobile-api-go/internal/controllers"
	"github.com/Eduardo-Brehm/mobile-api-go/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func SetupAuthRoutes(e *echo.Echo, authController *controllers.AuthController) {
	e.POST("/api/v1/autenticacao/cadastro", authController.Register)
	e.POST("/api/v1/autenticacao/login", authController.Login)
}

func SetupProfileRoutes(e *echo.Echo, profileController *controllers.ProfileController) {
	// Authenticated routes - require JWT token
	e.GET("/api/v1/perfil/eu", profileController.GetProfile, middlewares.AuthMiddleware())
	e.PATCH("/api/v1/perfil/eu", profileController.UpdateProfile, middlewares.AuthMiddleware())

	// Public route - no authentication needed
	e.GET("/api/v1/perfil/:username", profileController.GetPublicProfile)
}
