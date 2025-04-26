package main

import (
	"scaffold-demo/config"
	_ "scaffold-demo/config"
	"scaffold-demo/routers"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载程序配置
	// 2. 配置Gin
	r := gin.Default()
	logs.Info(nil, "server start")
	// 3. 注册路由
	routers.RegisterRouters(r)
	r.Run(config.Port)
}
