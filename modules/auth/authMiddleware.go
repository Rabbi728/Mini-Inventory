package auth

import (
	"mini-inventory/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, utils.ErrorResponse("Unauthorized", "Authorization header is required"))
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.JSON(401, utils.ErrorResponse("Unauthorized", "Bearer token is required"))
			c.Abort()
			return
		}

		authService := AuthService{}
		user, err := authService.GetUserByToken(token)
		if err != nil {
			c.JSON(401, utils.ErrorResponse("Unauthorized", err.Error()))
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Set("token", token)
		c.Next()
	}
}
