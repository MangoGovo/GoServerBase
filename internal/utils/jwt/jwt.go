package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"go-server-example/internal/exceptions"
	"go-server-example/pkg/config"

	"go.uber.org/zap"
	"time"
)

type UserClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

var secretKey string
var expireHour int

func init() {
	secretKey = config.Config.GetString("jwt.secret")
	expireHour = config.Config.GetInt("jwt.expireHour")
}

func GenerateJWT(userID uint) (string, error) {
	claims := UserClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireHour) * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                            // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                            // 生效时间
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(secretKey))

	return s, err
}

func ParseJwt(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		zap.L().Error("jwt解析失败", zap.Error(err))
		return nil, exceptions.ServerError
	}
	if userClaims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return userClaims, nil
	} else {
		return nil, exceptions.ServerError
	}
}
