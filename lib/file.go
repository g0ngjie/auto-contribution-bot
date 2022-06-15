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
func writeFile(writeDir string) {
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

	// 内容来源
	switch Cfg.ContentFrom {
	case 0:
		content := fmt.Sprintf("\n%s\n", Cfg.NewContent)
		write.WriteString(content)
		logrus.Infof("内容：%s", Cfg.NewContent)
	case 1:
		content := api.GetHitokoto()
		if content != nil {
			write.WriteString(fmt.Sprintf("\n### %s\n", content.Creator))
			write.WriteString(fmt.Sprintf("%s\n", content.Hitokoto))
			logrus.Infof("内容：%s", content.Hitokoto)
		}
	}

	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
