package main

import (
	"fmt"
	"time"
)

func init() { //先于main函数，谁前谁先执行
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

// 函数
func main() {
	//a := 15
	//fmt.Println("a:", a)
	//fmt.Println(mx(a))
	//fmt.Println(a)
	//
	//x, y := 1, 2
	//fmt.Println(x, y)
	//swap(&x, &y)
	//fmt.Println(x, y)
	//
	//half := func(x int) int { //函数赋值
	//	return x / 2
	//}
	//
	//fmt.Println(half)
	//fmt.Println(half(10))
	//
	//fmt.Println(cn(16, half))
	//fmt.Println(cn)

	var f = sum()
	fmt.Println(f, f(1, 3, 4))

	t1 := time.Now()
	var q = awaited(1)(1, 2, 3)
	t2 := time.Since(t1)
	fmt.Println(t2, q)

	fmt.Println(d1())
	fmt.Println(d2())
}

func mx(a int) int {
	a = 10
	return a
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func cn(x int, y func(i int) int) int { //函数作为参数
	return y(x)
}

func sum() func(args ...int) int { //匿名函数
	return func(args ...int) int {
		sum := 0
		for _, v := range args {
			sum += v
		}
		return sum
	}
}

func awaited(x int) func(...int) int { //延时函数
	time.Sleep(time.Duration(x) * time.Second)
	return sum()
}

func d1() int { //返回值无名，defer不修改返回值，只操作副本
	var i = 0
	defer func() {
		i = 2
		fmt.Println(i, "d1")
	}()
	return i
}

func d2() (i int) { //返回值有名，先执行return，后执行defer
	defer func() {
		i++
		fmt.Println(i, "d2")
	}()
	return 5
}
