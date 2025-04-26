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
	Port string
)

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
	// 从环境变量中读取配置
	viper.SetDefault("LOG_LEVEL", "debug")

	// 获取程序启动端口号配置
	viper.SetDefault("PORT", ":8080")

	viper.AutomaticEnv()
	// 获取配置信息
	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	// 初始化日志配置
	initLogConfig(logLevel)
}
