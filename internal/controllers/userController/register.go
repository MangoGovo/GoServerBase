package userController

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"go-server-example/internal/exceptions"
	"go-server-example/internal/services/userService"
	"go-server-example/internal/utils/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type registerReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register 注册
func Register(c *gin.Context) {
	var data registerReq
	if err := c.ShouldBindJSON(&data); err != nil {
		response.AbortWithException(c, exceptions.ParamsError, err)
	}

	user, err := userService.GetUserByUsername(data.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		response.AbortWithException(c, exceptions.UserExisted, fmt.Errorf("%s用户已经存在", user.Username))
		return
	}

	if err = userService.SaveUser(data.Username, data.Password); err != nil {
		response.AbortWithException(c, exceptions.ServerError, err)
		return
	}
	zap.L().Info("用户创建成功")
	response.JsonSuccess(c, nil)
}
