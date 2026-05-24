package main

//线程安全——加锁
import (
	"fmt"
	"sync"
)

var sum int = 0
var wait sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		sum++
		lock.Unlock()
	}
	wait.Done()
}

func sub() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		sum--
		lock.Unlock()
	}
	wait.Done()
}

func main() {
	cnt := 0
	for i := 0; i < 1000; i++ {
		sum = 0
		wait.Add(2)
		go add()
		go sub()
		wait.Wait()
		//fmt.Println(sum)
		if sum != 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
