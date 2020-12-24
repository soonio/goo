package basic

import (
	"fmt"
	"goo/tools"
)

// 这一部分内容主要关于go中的基本类型

// 定义常量
func defineConst() {
	tools.ShowPosition("定义常量")
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

	tools.ShowPosition("定义单个变量") // 这个函数无需理解
	var name string = "no name"
	name = "张三"
	fmt.Printf("变量定义 name=%v, 类型:%T \n", name, name)

	tools.ShowPosition("批量定义")
	var (
		sex      uint8
		age      uint8
		describe string
	)
	fmt.Printf("默认值 sex:%v age:%v describe:%v\n", sex, age, describe)

}

// 数据类型
func dataType() {

}

func Run(bit int) {
	fmt.Println(bit)
	if true {
		defineConst()
	}
	defineVar()
	dataType()
}
