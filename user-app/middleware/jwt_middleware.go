package middleware

import (
	"net/http"
	"strings"

	"go-lang/user-app/utils"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token is required"})
			c.Abort()
			return
		}

		// Bearer token 格式: "Bearer <token>"
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中，以便后续处理使用
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
