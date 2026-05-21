package main

import "fmt"

func main() {
	type Code int
	a := 10
	var c Code
	c = Code(a) //显示转换，自定义类型
	fmt.Println(a, c)

	type ResponseCode = int //起别名，不能创建新方法
	var r ResponseCode
	r = a //无需显示转换
	fmt.Println(a, r)
}
