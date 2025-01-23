package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-server-example/internal"
	"go-server-example/internal/midwares"
	"go-server-example/internal/utils/server"
	"go-server-example/pkg/config"
	_ "go-server-example/pkg/log"
)

func main() {
	// TODO Health Check

	// 如果配置文件中开启了调试模式
	if !config.Config.GetBool("server.debug") {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	internal.Init(r)

	server.Run(r, ":"+config.Config.GetString("server.port"))
}
