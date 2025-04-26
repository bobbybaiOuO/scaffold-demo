package config

import (
	"scaffold-demo/utils/logs"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	// TimeFormat defines the standard time format used throughout the application.
	TimeFormat string = "2006-01-02 15:04:05"
)

var (
	Port          string
	JwtSignKey    string
	JwtExpireTime int64 // JWT过期时间，单位分钟
	UserName      string
	Password      string
)

// ReturnData 定义返回数据的结构体
type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// NewReturnData 构造函数
func NewReturnData(status int, message string, data map[string]interface{}) *ReturnData {
	returnData := &ReturnData{}
	returnData.Status = 200
	returnData.Message = message
	data = make(map[string]interface{})
	returnData.Data = data
	return returnData
}

func initLogConfig(logLevel string) {
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	// 文件名和行号
	logrus.SetReportCaller(true)
	// 日志格式改为JSON格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: TimeFormat,
	})
}

func init() {
	logs.Debug(nil, "开始加载程序配置")
	// 配置日志等级
	viper.SetDefault("LOG_LEVEL", "debug")
	// 配置程序启动端口号
	viper.SetDefault("PORT", ":8080")
	// 配置JWT签名密钥
	viper.SetDefault("JWT_SIGN_KEY", "bobbybai")
	// 配置JWT过期时间
	viper.SetDefault("JWT_EXPIRE_TIME", 120)
	// 配置默认用户名和密码
	viper.SetDefault("USER_NAME", "bobby")
	viper.SetDefault("PASSWORD", "123456")

	viper.AutomaticEnv()
	// 获取配置信息
	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpireTime = viper.GetInt64("JWT_EXPIRE_TIME")
	UserName = viper.GetString("USER_NAME")
	Password = viper.GetString("PASSWORD")
	// 初始化日志配置
	initLogConfig(logLevel)
}
