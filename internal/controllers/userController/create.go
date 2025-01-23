package userController

import (
	"github.com/gin-gonic/gin"
	"go-server-example/internal/apiException"
	"go-server-example/internal/services/userService"
	"go-server-example/internal/utils"
	"go.uber.org/zap"
)

type createUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var data createUserReq
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.AbortWithException(c, apiException.ParamsError, err)
	}

	err := userService.CreateUser(data.Username, data.Password)
	if err != nil {
		utils.AbortWithException(c, apiException.ServerError, err)
	}
	zap.L().Info("用户创建成功")
	utils.JsonSuccessResponse(c, nil)
}
