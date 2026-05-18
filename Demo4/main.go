package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a%2 == 1 {
		fmt.Println("Yes")
	}

	switch a {
	case 0:
		fmt.Println("No")
		fallthrough
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("a>1")
	}

	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Println("Nil")
	default:
		fmt.Printf("%T\n", i)
	}
}
