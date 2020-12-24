package tools

import (
	"errors"
	"os"
	"strconv"
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
func Parse(args []string) (string, int) {
	bit := 0
	if len(args) == 2 {
		if IsNumeric(args[1]) {
			bit, _ = strconv.Atoi(args[1])
		}
	}
	return args[0], bit
}
