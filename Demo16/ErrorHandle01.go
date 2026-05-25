package main

import (
	"errors"
	"fmt"
)

//异常处理

func div(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return
	}
	return a / b, nil
}

func serve(a, b int) (res int, err error) {
	res, err = div(a, b)
	if err != nil {
		return //向上抛
	}
	res++
	return
}

func main() {
	num, err := serve(6, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}
