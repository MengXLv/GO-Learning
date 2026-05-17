package main

import "fmt"

func main() {
	var nums = [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums[5] = 100000
	fmt.Println(nums, len(nums), cap(nums))

	for i := 0; i < len(nums); i++ {
		nums[i] += 100
	}
	fmt.Println(nums)

	var nums2 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			nums2[i][j] = i*j + 10
		}
	}
	fmt.Println(nums2)

	var nums3 [][]string
	row1 := []string{"aaa"}
	row2 := []string{"bbb", "528"}
	row3 := []string{"ccc", "sdaff"}
	nums3 = append(nums3, row1)
	nums3 = append(nums3, row2)
	nums3 = append(nums3, row3)
	for i := range len(nums3) {
		fmt.Println(nums3[i], len(nums3[i]))
	}
	nums3[0][0] = "a"
	fmt.Println(nums3)
	nums3[0] = []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(nums3)
	fmt.Println("--------------------------------")
	for _, row := range nums3 {
		fmt.Println(row, "45")
		for _, col := range row {
			fmt.Println(col, len(col))
		}
	}
	fmt.Println("--------------------------------")

	n1 := [2][2]int{{1, 2}, {1, 6}}
	n2 := [2][2]int{{1, 2}, {1, 6}}
	fmt.Println("n1==n2", n1 == n2)

	n3 := [][]int{{1, 2}, {1, 6}}
	fmt.Printf("%T\n", n3)
	cur := append(n3[0], 3)
	fmt.Print(n3, cur)
}
