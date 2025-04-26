package middlewares

import (
	"scaffold-demo/config"
	"scaffold-demo/utils/jwtutil"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

// JwtAuth 中间件，用于验证JWT
func JwtAuth(c *gin.Context) {
	// 1. 除了登录和登出接口，其他接口都需要验证JWT
	requestURL := c.FullPath()
	logs.Debug(map[string]interface{}{
		"requestURL": requestURL,
	}, "")
	if requestURL == "/api/auth/login" || requestURL == "/api/auth/logout" {
		logs.Debug(map[string]interface{}{
			"requestURL": requestURL,
		}, "登录和退出不需要验证JWT")
		c.Next()
		return
	}
	returnData := config.NewReturnData(200, "success", nil)
	// 2. 验证JWT
	// tokenStr := c.GetHeader("Authorization")
	tokenStr := c.Request.Header.Get("Authorization")
	// token为空
	if tokenStr == "" {
		// 请求没有携带token，返回401状态码
		returnData.Status = 401
		returnData.Message = "请求没有携带token,请登录后尝试"
		c.JSON(200, returnData)
		c.Abort()
		return
	}
	// token不为空，验证token
	claims, err := jwtutil.ParseToken(tokenStr)
	if err != nil {
		// token验证失败，返回401状态码
		returnData.Status = 401
		returnData.Message = "token验证为通过"
		c.JSON(200, returnData)
		c.Abort()
		return
	}
	// token验证通过，将claims存储到上下文中
	c.Set("claims", claims)
	c.Next()
}
