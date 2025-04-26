package routers

import (
	"scaffold-demo/routers/auth"

	"github.com/gin-gonic/gin"
)

// RegisterRouters 注册路由
func RegisterRouters(r *gin.Engine) {
	// 登录路由配置
	// 1. 登录：/login 登录接口
	// 2. 退出：/logout 退出接口
	// 3.      /api/auth/login 登录接口
	//         /api/auth/logout 退出接口
	apiGroup := r.Group("/api")
	auth.RegisterSubRouters(apiGroup)
	
}
