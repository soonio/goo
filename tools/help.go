package tools

import (
	"log"
	"runtime"
)

func ShowPosition() {
	skip := 1
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		line--
		log.Printf("pc:%v, file:%s, line:%d, ok:%v skip:%d \n\n", pc, file, line, ok, skip)
	}
}
