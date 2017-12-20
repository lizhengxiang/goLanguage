package main

import (
	"fmt"
	"os"
)

func main() {
	//变量在申明的时候会初始化
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
