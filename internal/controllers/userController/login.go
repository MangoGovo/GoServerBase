package userController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-server-example/internal/exceptions"
	"go-server-example/internal/services/userService"
	"go-server-example/internal/utils/jwt"
	"go-server-example/internal/utils/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type loginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResp struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var data loginReq
	var resp loginResp
	if err := c.ShouldBindJSON(&data); err != nil {
		response.AbortWithException(c, exceptions.ParamsError, err)
	}

	user, err := userService.GetUserByUsername(data.Username)
	if err != nil {
		response.AbortWithException(c, exceptions.WrongPasswordOrUsername, err)
	}

	if user.Password != data.Password || errors.Is(err, gorm.ErrRecordNotFound) {
		response.AbortWithException(c, exceptions.WrongPasswordOrUsername, err)
		return
	} else if err != nil {
		response.AbortWithException(c, exceptions.ServerError, err)
		return
	}

	token, err := jwt.GenerateJWT(user.ID)
	resp.Token = token
	if err != nil {
		response.AbortWithException(c, exceptions.ServerError, err)
		return
	}

	zap.L().Info("用户" + user.Username + "登陆成功")
	response.JsonSuccess(c, loginResp{
		Token: token,
	})
}
