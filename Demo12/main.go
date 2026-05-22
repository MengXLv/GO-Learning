package main

import "fmt"

// 接口
type pet interface {
	eat()
	sleep()
	drink()
}
type em interface { //空接口可以接受任意类型
}

type cat struct {
	name string
}

type dog struct {
	name string
}

func (c *cat) eat() {
	fmt.Println("cat is eating")
}
func (c *cat) sleep() {
	fmt.Println("cat is sleeping")
}
func (c *cat) drink() {
	fmt.Println("cat is drinking")
}
func (d *dog) eat() {
	fmt.Println("dog is eating")
}
func (d *dog) sleep() {
	fmt.Println("dog is sleeping")
}
func (d *dog) drink() {
	fmt.Println("dog is drinking")
}

func main() {
	var mimi *cat
	mimi = &cat{"mimi"}
	mimi.sleep()
	mimi.drink()
	mimi.eat()
	var p pet //实现了对应接口的一个变量
	p = mimi
	fmt.Println(p)
	ca, ok := p.(*cat) //类型断言
	fmt.Println(ca, ok)
	do, ok := p.(*dog)
	fmt.Println(do, ok)

	var e em
	e = mimi
	fmt.Println(e)

}
