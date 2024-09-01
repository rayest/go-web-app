package main

import (

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go-lang/user-app/middleware"
	"go-lang/user-app/routes"
)

type User struct {
	ID       string `gorm:"column:id" json:"id"`
	Username string `gorm:"column:user_name" json:"username"`
	Mobile   string `gorm:"column:mobile" json:"mobile"`
}

// 指定表名为 "t_c_user"
func (User) TableName() string {
	return "t_user"
}

var db *gorm.DB

func init() {
	var err error
	// 连接到 MySQL 数据库
	dsn := "root:Rayest@1108@(127.0.0.1:3306)/go_app?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate database schema")
	}
}

func main() {
	r := gin.Default()

	// 注册无需校验的登录路由
	routes.AuthRoutes(r)

	// 应用 JWT 中间件（需要保护的路由）
	protected := r.Group("/api/v1")
	protected.Use(middleware.JWTMiddleware())

	// 注册需要 JWT 校验的用户路由
	routes.UserRoutes(protected)

	// 启动服务
	r.Run(":18000")

}
