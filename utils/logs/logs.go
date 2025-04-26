package logs

import "github.com/sirupsen/logrus"

// Debug 级别日志输出，用于调试程序
func Debug(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Debug(msg)
}

// Info 级别日志输出，用于记录程序运行信息
func Info(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Info(msg)
}

// Warning 级别日志输出，用于记录程序运行警告信息
func Warning(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Warning(msg)
}

// Error 级别日志输出，用于记录程序运行错误信息
func Error(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Error(msg)
}

