package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num [3]int = [3]int{1, 2, '李'}
	fmt.Println(num, unsafe.Sizeof(num), unsafe.Sizeof(num[0]))

	var num1 = []int{2, 4, 6}
	fmt.Println(num1)

	num2 := []float64{2.1, 3.55}
	fmt.Println(num2)

	const (
		a = 6
		b = 7
		c = "ssdf"
	)
	fmt.Println(a, b, c)
	//a = 9
	//fmt.Println(a)
	//cannot assign to a (neither addressable nor a map index expression)

	const (
		i = iota
		j
		k
		l = "89"
		m = iota
		n = iota
	)
	fmt.Println(i, j, k, l, m, n)
}
