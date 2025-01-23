package userService

import (
	"go-server-example/internal/models"
	"go-server-example/pkg/database"
)

// CreateUser 创建用户
func CreateUser(username, password string) error {
	user := &models.User{
		Username: username,
		Password: password,
	}
	res := database.DB.Create(&user)
	return res.Error
}
