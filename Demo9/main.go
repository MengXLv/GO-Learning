package main

import "fmt"

func main() {
	var usemap map[int]string = map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}
	fmt.Println(usemap)
	fmt.Printf("%#v\n", usemap[5]) //不存在返回默认值

	val, ok := usemap[1]
	fmt.Println(val, ok)

	val, ok = usemap[6]
	fmt.Println(val, ok)
}
