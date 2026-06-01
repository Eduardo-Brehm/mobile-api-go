package middlewares

import (
	"net/http"

	"github.com/Eduardo-Brehm/mobile-api-go/internal/utils"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// get the auth header from the request
			authHeader := c.Request().Header.Get("Authorization")

			// check if the auth header is present and starts with "Bearer "
			if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "token não fornecido"})
			}
			token := authHeader[7:]

			// verify the token and get the userID
			userID, err := utils.VerifyToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "token inválido"})
			}

			// store the userID in the context for use in the next handlers
			c.Set("userId", userID)

			return next(c)
		}
	}
}
