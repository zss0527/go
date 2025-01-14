package main

import (
	"gin_gorm_demo/internal/config"
	"gin_gorm_demo/internal/controller"
	"gin_gorm_demo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	conf := config.LoadConfig()

	// 初始化 Gin 引擎
	r := gin.Default()

	// 初始化用户控制器
	userCtrl := controller.NewUserController()

	// 注册用户路由
	router.RegisterUserRoutes(r, userCtrl)

	// 启动服务
	r.Run(conf.Server.Address)
}
