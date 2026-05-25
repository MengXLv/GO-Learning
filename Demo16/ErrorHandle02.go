package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	res, err := os.ReadFile("123") //中断处理，常用于初始化
	if err != nil {
		log.Fatal(err) //打印错误日志
		//panic(err) //打印错误堆栈
	}
	fmt.Println(string(res))
}
func main() {
	fmt.Println("Hello World")
}
