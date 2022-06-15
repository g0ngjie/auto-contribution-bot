package lib

import (
	"fmt"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Welcome() {
	fmt.Println("Welcome to the bot!")
}

// 配置一个日志输出
type AppHook struct {
	AppName string
}

func (h *AppHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = h.AppName
	return nil
}

func init() {
	// 日志级别 debug 或 info
	logrus.SetLevel(logrus.DebugLevel)
	// 日志追踪
	logrus.SetReportCaller(false)
	// 添加钩子
	h := &AppHook{AppName: "BOT"}
	logrus.AddHook(h)
	// 设置日志格式
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: time.RFC3339,
	})
	// 设置日志输出
	writer := &lumberjack.Logger{
		// 日志文件路径
		Filename: "bot.log",
		// 日志文件大小，超过后会自动生成新的日志文件
		MaxSize: 100, // megabytes
		// 日志文件最多保存备份个数
		MaxBackups: 7,
		// 日志文件最多保存多少天
		MaxAge: 7, //days
		// 日志文件是否压缩
		Compress: true,
	}
	logrus.SetOutput(writer)
}
