package main

import (
	"fmt"
	"os"
)

func main() {
	//变量在申明的时候会初始化
	s, sep := "", ""
	// range 每次迭代产生一对值,索引以及该索引的值
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
