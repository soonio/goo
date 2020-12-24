package basic

import (
	"fmt"
	"goo/tools"
)

// 这一部分内容主要关于go中的基本类型

// 定义常量
func defineConst() {
	const PI = 3.14159       // 隐式类型定义
	const Name string = "清流" // 显式类型定义 ，指出常量的类型名称

	fmt.Printf("隐式定义常量，PI:%v, 显式定义常量Name: %v\n", PI, Name)

	// 支持更骚的操作，并行赋值。
	const A, B, C = 1, 2, 3
	fmt.Printf("并行赋值，常量A: %v B: %v C: %v\n", A, B, C)

	const (
		a = iota
		b // = iota
		c // = iota
	)
	// iota 每次遇到const会重置为0
	fmt.Printf("iota增量赋值，常量a: %v b: %v c: %v\n", a, b, c)
}

// 定义变量
func defineVar() {
	//var identifier type
	var name string
	name = "张三"
	fmt.Printf("变量定义 name=%v, 类型:%T \n", name, name)
	tools.ShowPosition() // 这个函数无需理解

}

func Run() {
	defineConst()
	defineVar()
}
