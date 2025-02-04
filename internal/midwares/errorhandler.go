package midwares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-server-example/internal/exceptions"
	"go-server-example/internal/utils/response"
	"go.uber.org/zap"
)

// ErrHandler 中间件用于处理请求错误。
// 如果存在错误，将返回相应的 JSON 响应。
func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 向下执行请求
		c.Next()

		// 如果存在错误，则处理错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			if err != nil {
				var apiErr *exceptions.Error

				// 尝试将错误转换为 exceptions
				ok := errors.As(err, &apiErr)

				// 如果转换失败，则使用 ServerError
				if !ok {
					apiErr = exceptions.ServerError
					zap.L().Error("遇到了未知的异常:", zap.Error(err))
				}

				response.JsonError(c, apiErr.Code, apiErr.Msg)
				return
			}
		}
	}
}

// HandleNotFound 处理 404 错误。
func HandleNotFound(c *gin.Context) {
	err := exceptions.NotFound
	// 记录 404 错误日志
	zap.L().Warn("404 Not Found",
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
	)
	response.Json(c, http.StatusNotFound, err.Code, err.Msg, nil)
}
