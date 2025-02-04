package userService

import (
	"go-server-example/internal/models"
	"go-server-example/pkg/database"
)

// SaveUser 创建用户
func SaveUser(username, password string) error {
	user := &models.User{
		Username: username,
		Password: password,
	}
	res := database.DB.Create(&user)
	return res.Error
}
