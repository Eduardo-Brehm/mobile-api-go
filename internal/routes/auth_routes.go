package routes

import (
	"github.com/Eduardo-Brehm/mobile-api-go/internal/controllers"
	"github.com/labstack/echo/v4"
)

func SetupAuthRoutes(e *echo.Echo, authController *controllers.AuthController) {
	e.POST("/api/v1/autenticacao/cadastro", authController.Register)
	e.POST("/api/v1/autenticacao/login", authController.Login)
}
