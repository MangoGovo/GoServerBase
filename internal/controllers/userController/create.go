package userController

import (
	"github.com/gin-gonic/gin"
	"go-server-example/internal/exceptions"
	"go-server-example/internal/services/userService"
	"go-server-example/internal/utils/response"
	"go.uber.org/zap"
)

type createUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var data createUserReq
	if err := c.ShouldBindJSON(&data); err != nil {
		response.AbortWithException(c, exceptions.ParamsError, err)
	}

	err := userService.CreateUser(data.Username, data.Password)
	if err != nil {
		response.AbortWithException(c, exceptions.ServerError, err)
	}
	zap.L().Info("用户创建成功")
	response.JsonSuccess(c, nil)
}
