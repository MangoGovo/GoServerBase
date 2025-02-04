package userService

import (
	"go-server-example/internal/models"
	"go-server-example/pkg/database"
)

// GetUserByUsername 通过用户名获取用户
func GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	res := database.DB.Where("username = ?", username).First(user)
	return user, res.Error
}
