package main

import (
	"fmt"
	"regexp"
)

func main() {
	//验证邮箱
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[A-Za-z0-9-]+\.[A-Za-z]{2,}`)
	fmt.Println(re.MatchString("22@whu.com"), re.MatchString("@"))

	//验证手机号
	phone := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	fmt.Println(phone.MatchString("123"), phone.MatchString("0518-22-44"))
}
