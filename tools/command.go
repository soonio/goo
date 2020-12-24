package tools

import (
	"errors"
	"os"
)

//关于命令行内容处理的工具方法
func GetArgs() ([]string, error) {

	l := len(os.Args)
	if l == 1 {
		return make([]string, 0), errors.New("请至少包含一个命令参数")
	} else {
		return os.Args[1:l], nil
	}
}

// 简单解析命令
func Parse(args []string) map[string]string {
	command := make(map[string]string, 5)
	command["command"] = args[0]
	return command
}
