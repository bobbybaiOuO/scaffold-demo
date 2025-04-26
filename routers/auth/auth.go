package auth

import (
	"scaffold-demo/controllers/auth"

	"github.com/gin-gonic/gin"
)

// RegisterSubRouters 注册子路由
func RegisterSubRouters(r *gin.RouterGroup) {
	// 配置登录功能的路由策略
	authGroup := r.Group("/auth")
	// 登录功能
	login(authGroup)
	// 退出功能
	logout(authGroup)

}

// login 实现登录接口
func login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
}

// logout 实现退出接口
func logout(authGroup *gin.RouterGroup) {
	authGroup.GET("/logout", auth.Logout)
}
