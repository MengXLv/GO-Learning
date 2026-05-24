package main

//线程不安全
import (
	"fmt"
	"sync"
)

var sum int = 0
var wait sync.WaitGroup

func add() {
	for i := 0; i < 1000; i++ {
		sum++
	}
	wait.Done()
}

func sub() {
	for i := 0; i < 1000; i++ {
		sum--
	}
	wait.Done()
}

func main() {
	//wait.Add(1)
	cnt := 0
	for i := 0; i < 1000; i++ {
		sum = 0
		wait.Add(2)
		go add()
		go sub()
		wait.Wait()
		fmt.Println(sum)
		if sum != 0 {
			cnt++
		}
	}
	//wait.Wait()
	fmt.Println(cnt)
}
