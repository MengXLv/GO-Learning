package main

import (
	"fmt"
	"strconv"
)

//学习类型转换
//不支持隐型转换

func main() {
	var a int = 42
	var b float32 = 6.0
	//fmt.Println(a / b)//无法转换
	c := float32(a) / b
	fmt.Println(c)
	fmt.Printf("c is of type %T\n", c)

	var s string = "6"
	d, e := strconv.Atoi(s)
	fmt.Printf("d is of type %T\n", d)
	fmt.Println(e)

	var num int = 100
	p := strconv.Itoa(num)
	fmt.Printf("p is of type %T\n", p)

	s = "6.25"
	f, e := strconv.ParseFloat(s, 32) //无论参数为何，都是float64，第二个参数决定内部解析
	fmt.Printf("f is of type %T\n", f)

	f = 24.2665
	s = strconv.FormatFloat(f, 'E', -1, 32)
	fmt.Printf("s is of type %T\n", s)
	fmt.Println(s)

	//var o float64 = 1.1
	//var q float64 = 1.2
	//fmt.Println(o + q)
}
