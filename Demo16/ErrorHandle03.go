package main

import (
	"fmt"
	"runtime/debug"
)

func read() {
	defer func() { //使用defer捕获异常
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
		}
	}()
	var list = []int{1, 2}
	fmt.Println(list[2])
}

func main() {
	read()
	fmt.Println("运行正常")
}
