package lib

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"bot/api"

	"github.com/sirupsen/logrus"
)

const (
	fileName = "bot"
)

// 获取日期
func getFmtDate() string {
	// 获取当前时间
	now := time.Now()
	// 格式化 YYYY-MM-DD
	return now.Format("2006-01-02")
}

// 写入内容
func writeFile(writeDir string) *string {
	filePath := fmt.Sprintf("%s/%s.%s", writeDir, fileName, Cfg.FileType)
	var (
		file *os.File
		err  error
	)

	// 判断文件是否需要新增
	if Cfg.NewFile {
		filePath = fmt.Sprintf("%s/%s-%s.%s", writeDir, fileName, getFmtDate(), Cfg.FileType)
		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		logrus.Infof("新增文件：%s", filePath)
	} else {
		// 判断文件是否存在
		if _, err = os.Stat(filePath); err != nil {
			file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
			logrus.Infof("新增文件：%s", filePath)
		} else {
			file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
			logrus.Infof("追加文件：%s", filePath)
		}
	}
	if err != nil {
		logrus.Infof("文件打开失败：%s", err.Error())
	}

	//及时关闭file句柄
	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	var content string
	// 内容来源
	switch Cfg.ContentFrom {
	case 0:
		content = fmt.Sprintf("\n%s\n", Cfg.NewContent)
		write.WriteString(content)
		logrus.Infof("内容：%s", Cfg.NewContent)
	case 1:
		contentBody := api.GetHitokoto()
		if contentBody != nil {
			content = fmt.Sprintf("\n%s\n", contentBody.Hitokoto)
			write.WriteString(fmt.Sprintf("\n### %s\n", contentBody.Creator))
			write.WriteString(fmt.Sprintf("%s\n", contentBody.Hitokoto))
			logrus.Infof("内容：%s", contentBody.Hitokoto)
		}
		content = contentBody.Hitokoto
	}

	//Flush将缓存的文件真正写入到文件中
	write.Flush()

	// 如果content内容过长，则截取 + ...
	// 最长限制在30字符
	var maxLen int = 30
	if len(content) > maxLen {
		content = content[:maxLen] + "..."
	}
	return &content
}
