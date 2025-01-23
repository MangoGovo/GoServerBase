package internal

import (
	"github.com/gin-gonic/gin"
	"go-server-example/internal/controllers/userController"
)

// Init 初始化路由
func Init(r *gin.Engine) {
	r.POST("/user", userController.CreateUser)
}
