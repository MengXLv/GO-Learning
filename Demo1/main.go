package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	s := "秀米"
	fmt.Printf("nihao %s\n", s)
	fmt.Printf("我是%s\n", "刘邦")
	fmt.Printf("%d\n", 3)
	fmt.Printf("%f\n", 4.0)
	fmt.Printf("%.2f\n", 4.0)
	fmt.Printf("%T %T\n", 4.0, 9)
	fmt.Printf("%v\n", "")
	fmt.Printf("%v\n", 'a'-20)

	//var temp string
	//fmt.Scan(&temp)
	//fmt.Printf("%s\n", temp)

	var a int = 1
	var b uint = 2
	var c uint8 = 3
	var d byte = 'a'
	fmt.Println(a, b, c, d)

	var e rune = '中'
	var f byte = '6'
	fmt.Println(e, f)

	var g bool = true
	fmt.Println(g, g && false, g || false)

	var h float64 = 3.14
	fmt.Println(h)
}
