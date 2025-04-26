package auth

import (
	"scaffold-demo/config"
	"scaffold-demo/utils/jwtutil"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

// UserInfo 用户信息
type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login 登录
func Login(c *gin.Context) {
	// 1. 从请求中获取用户名和密码
	userInfo := UserInfo{}
	returnData := config.NewReturnData(200, "success", nil)
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		returnData.Status = 401
		returnData.Message = err.Error()
		c.JSON(200, returnData)
		return
	}
	logs.Debug(map[string]interface{}{"userInfo": userInfo.Username, "password": userInfo.Password}, "开始验证登录信息")
	// 2. 验证用户名和密码是否正确
	if userInfo.Username == config.UserName && userInfo.Password == config.Password {
		// 3. 生成JWT
		token, err := jwtutil.GenerateToken(userInfo.Username)
		if err != nil {
			logs.Error(map[string]interface{}{"userInfo": userInfo.Username, "err_msg": err.Error()}, "用户名密码正确但是生成Token失败")
			c.JSON(200, gin.H{
				"message": "生成Jwt Token失败",
				"status":  401,
			})
			return
		}
		// token 生成成功，返回给客户端
		logs.Info(map[string]interface{}{"userInfo": userInfo.Username}, "用户名密码正确,登录成功")
		returnData.Message = "登录成功"
		returnData.Data["token"] = token
		c.JSON(200, returnData)
		return
	} else {
		// 用户名或密码错误
		logs.Error(map[string]interface{}{"userInfo": userInfo.Username}, "用户名或密码错误")
		c.JSON(200, gin.H{
			"message": "用户名或密码错误",
			"status":  401,
		})
		return
	}

}

// Logout 退出
func Logout(c *gin.Context) {
	// 退出
	c.JSON(200, gin.H{
		"message": "退出成功",
		"status":  200,
	})
	logs.Debug(nil, "用户已退出")
}
