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

type Stringer interface { //自定义约束，该
	String() string
}
type Person struct {
	name string
}

func (p Person) String() string {
	return p.name
}

func stringPrint[T Stringer](s T) { //该s必须实现Stringer接口
	fmt.Println(s.String())
}

func main() {
	a, b := 1.0, 2.0
	fmt.Println(plus(a, b))
	MyPrintf("588")
	fmt.Println(equals(1.0, 2.0))
	var p Person
	p.name = "Jack"
	stringPrint(p)
}
