package main

// golang学习 开发环境
//	- golang idea
//	- log version go1.15.5 darwin/amd64

// 使用import 加载多个包
import (
	"fmt"
	"goo/example/basic"
	"goo/example/file"
	"goo/example/orm"
	"goo/tools"
)

// 上面的fmt是系统包
// goo/* 是自定义包，需要使用 log mod init goo 生成 log.mod 才能正常引入

// main 方法是go规定的入口方法，命令行中执行 log run main.log 默认会执行main方法
func main() {
	args, err := tools.GetArgs() // 调用自定义的tools包 获取命令行输入的参数
	if err != nil {              // 简单的判断是否返回错误如果返回错误，
		fmt.Printf("%v.\n如下命令格式:\n\tlog run main.log packagename.funcitonName\n", err)
		return
	}
	command, bit := tools.Parse(args) // 解析命令行输入的参数

	// 本示例为了简单，使用简单的逻辑判断，确定是否执行对应的demo方法
	switch command {
	case "basic":
		basic.Run(bit)
		break
	case "file":
		file.Run()
		break
	case "orm":
		orm.Run()
		break
	case "test":
		break
	default:
		fmt.Println("没有该命令有效的方法")
	}
}
