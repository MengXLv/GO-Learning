package main

import (
	"fmt"
	"sort"
)

//切片

func main() {
	var nameList []string //可变数组
	fmt.Printf("%T", nameList)
	nameList = append(nameList, "aaa")
	nameList = append(nameList, "bbb")
	fmt.Println(nameList[0])

	ageList := make([]int, 3) //生成固定长度切片
	fmt.Printf("%T", ageList)
	fmt.Println(ageList)
	ageList[0] = 8
	ageList[1] = 10
	ageList[2] = 9
	fmt.Println(ageList)

	sort.Ints(ageList) //切片排序
	fmt.Println(ageList)
	sort.Sort(sort.Reverse(sort.IntSlice(ageList)))
	fmt.Println(ageList)

	nums := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums)
	slice := nums[1:3] //从数组切分，左闭右开
	fmt.Println(slice)
}
