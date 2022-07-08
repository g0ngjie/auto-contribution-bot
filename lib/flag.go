package lib

import (
	"flag"
	"fmt"
	"strings"
)

// 获取toml配置文件
func FlagArgs() {
	// 默认配置文件路径
	defaultName := "conf.toml"
	flag.StringVar(&tomlPath, "c", "conf.toml", "config file path")
	flag.Parse()
	// git-repo目录命名
	if defaultName != tomlPath {
		// 截取配置文件名
		name := tomlPath[:strings.Index(tomlPath, ".")]
		gitDir += fmt.Sprintf("-%s", name)
	}
}
