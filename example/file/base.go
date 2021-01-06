package file

import (
	"fmt"
	"os"
)

func Run() {
	file, err := os.Open("./example/test.txt") // 填写相对于main.go文件的地址
	if err != nil {
		fmt.Println(err)
		fmt.Println("打开文件失败")
		return
	}
	defer file.Close()

	tmp := make([]byte, 128)
	n, err := file.Read(tmp)
	if err != nil {
		fmt.Println("读取文件失败")
		return
	}

	fmt.Println(n)
	fmt.Println(tmp)

}
