package main

import "fmt"

func main() {
	a := 15
	fmt.Println("a:", a)
	fmt.Println(mx(a))
	fmt.Println(a)

	x, y := 1, 2
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)

	half := func(x int) int { //函数赋值
		return x / 2
	}

	fmt.Println(half)
	fmt.Println(half(10))

	fmt.Println(cn(16, half))
	fmt.Println(cn)

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
