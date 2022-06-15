package lib

import (
	"log"

	"github.com/BurntSushi/toml"
)

// 获取toml配置文件

type Config struct {
	NewFile     bool   `toml:"new_file"`     // 是否每天生成新的文件，默认为False 非新增，则在执行文件中添加新的行
	FileType    string `toml:"file_type"`    // 文件类型，默认为md
	ContentFrom uint8  `toml:"content_from"` // 内容来源，0：默认读取new_content 1：一言，默认为1
	NewContent  string `toml:"new_content"`  // 新增内容，默认为空
	GitUrl      string `toml:"git_url"`      // git地址
	GitUser     string `toml:"git_user"`     // git用户名
	GitPass     string `toml:"git_pass"`     // git密码
}

var Cfg Config

func init() {
	_, err := toml.DecodeFile("conf.toml", &Cfg)
	if err != nil {
		log.Fatal(err)
	}
	if Cfg.NewContent == "" {
		Cfg.NewContent = getFmtDate()
	}
	if Cfg.FileType == "" {
		Cfg.FileType = "md"
	}
}
