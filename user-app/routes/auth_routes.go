package routes

import (
	"go-lang/user-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// 登录路由
func AuthRoutes(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		var loginRequest LoginRequest
		if err := c.BindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
			return
		}

		// 验证用户凭据（此处为示例，实际应用中应从数据库查询）
		if loginRequest.Login == "rayest" && loginRequest.Password == "password" {
			token, err := utils.GenerateToken(loginRequest.Login)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Error generating token"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token})
			return
		}

		// 用户凭据无效时返回 400 Bad Request
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials"})
	})
}
