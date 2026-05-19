package main

import "fmt"

func main() {
	var ip *int
	fmt.Printf("%x\n", ip)
	//*ip= 2 //地址复制，错误
	fmt.Println(ip)
	ip = new(5)
	fmt.Printf("%x %d\n", ip, *ip)

	a := []int{1, 2, 3} //指针数组
	var iptr [3]*int
	for i := 0; i < len(a); i++ {
		iptr[i] = &a[i]
	}
	for i := 0; i < len(a); i++ { //随机分配
		fmt.Printf("%d\t%d\n", *iptr[i], iptr[i])
	}

	s := []string{"asasd", "a", "csdasd"}
	var iptrs [3]*string
	for i := 0; i < len(s); i++ {
		iptrs[i] = &s[i]
		fmt.Printf("%s\t%d\n", *iptrs[i], iptrs[i]) //16字节
	}
}
