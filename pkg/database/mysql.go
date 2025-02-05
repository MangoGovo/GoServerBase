package database

import (
	"fmt"

	"go-server-example/pkg/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 是全局数据库连接实例
var DB *gorm.DB

func init() {
	// 从配置中获取数据库连接所需的参数
	user := config.Config.GetString("database.user") // 数据库用户名
	pass := config.Config.GetString("database.pass") // 数据库密码
	port := config.Config.GetString("database.port") // 数据库端口
	host := config.Config.GetString("database.host") // 数据库主机
	name := config.Config.GetString("database.name") // 数据库名称

	// 构建数据源名称 (DSN)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	// 使用 GORM 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭外键约束以提升迁移速度
	})

	// 如果连接失败，返回错误
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	// 自动迁移数据库结构
	if err = autoMigrate(db); err != nil {
		zap.L().Fatal(err.Error())
	}

	// 将数据库实例赋值给全局变量 DB
	DB = db
	zap.L().Info("数据库初始化成功")
}
