package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	rg.GET("/users", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		c.JSON(http.StatusOK, gin.H{"message": "Success", "user_id": userID})
	})

	// 其他用户相关的路由
}
