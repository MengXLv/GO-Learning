package main

import "fmt"

// 泛型
type Number interface { //接口定义
	int | float64 | uint | byte
}

func plus[T Number](a, b T) T {
	return a + b
}

func MyPrintf[T any](a T) { //any任意约束
	fmt.Println(a)
}

func equals[T comparable](a T, b T) bool { //comparable约束
	return a == b
}

func main() {
	a, b := 1.0, 2.0
	fmt.Println(plus(a, b))
	MyPrintf("588")
	fmt.Println(equals(1.0, 2.0))
}
