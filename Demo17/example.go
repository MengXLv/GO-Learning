package main

import "fmt"

func main() {
	fmt.Println(swap('a', 'b'))
	fmt.Println(Contains([]string{"asd"}, "a"))
	fmt.Println(Union([]int{1, 2, 3, 2, 3}))
}

func swap[T any](a, b T) (T, T) { //交换
	return b, a
}

func Contains[T comparable](s []T, e T) bool { //分析切片是否包含某个元素
	if len(s) == 0 {
		return false
	}
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func Union[T comparable](s []T) []T { //去重
	var ans []T
	if len(s) == 0 {
		return ans
	}
	m := make(map[T]bool)
	for _, v := range s {
		if !m[v] {
			m[v] = true
			ans = append(ans, v)
		}
	}
	return ans
}
