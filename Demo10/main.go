package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
}

type Employee struct { //组合
	Person
	Name string //冲突优先本类
	Id   string
}

//	type Student struct {
//		Name   string
//		Age    int
//		Gender string
//	}
//
//	func (s Student) Study() {
//		fmt.Println(s.Name, "is studying")
//	}

func (emp Employee) SetId(id string) { //无效修改，值传递
	emp.Id = id
}
func (emp *Employee) SetId2(id string) { //有效修改，引用传递
	emp.Id = id
}

type Animal struct {
	Name   string `json:"name"` //结构体tag，进行类型转换
	Age    int    `json:"age"`
	Keeper string `json:"-"` //无视
}

func main() {
	//s := Student{
	//	Name:   "LiMing",
	//	Age:    15,
	//	Gender: "Female",
	//}
	//s.Study()

	e := new(Employee)
	e.Id, e.Name, e.Person.Name = "001", "Peter", "张三"
	fmt.Println(e)
	e.SetId("002")
	fmt.Println(e.Id)
	e.SetId2("002")
	fmt.Println(e.Id)

	a := Animal{"Cat", 15, "刘邦"}
	byteData, _ := json.Marshal(a)
	fmt.Println(string(byteData))
}
