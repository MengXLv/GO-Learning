package main

import (
	"fmt"
	"sync"
	"time"
)

//	func say() {
//		for i := 0; i < 3; i++ {
//			fmt.Printf("%d ", i)
//			time.Sleep(time.Second)
//		}
//	}

func sum(s []int, c chan int) {
	r := 0
	for _, x := range s {
		r += x
	}
	c <- r //sum传到通道c
}

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) //关闭通道
}

func fb(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println("x", x, "y", y)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func worker(id string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker " + id + " is started")
	fmt.Println("worker " + id + " is done")
}

var done = make(chan bool)

func see(a int) {
	fmt.Println("开始")
	time.Sleep(time.Duration(a) * time.Second)
	fmt.Println("结束")
	done <- true
}

func main() {
	//go say() //启动Goroutine
	//for i := 0; i < 3; i++ {
	//	fmt.Println("hello world")
	//	time.Sleep(time.Second)
	//}

	//go fmt.Println("hello world") //可能没有输出
	//time.Sleep(1 * time.Second)

	//ch := make(chan int)
	//arr := []int{1, 2, 3, 4, 5}
	//go sum(arr[:len(arr)/2], ch)
	//go sum(arr[len(arr)/2:], ch)
	//x, y := <-ch, <-ch //随机
	//fmt.Println(x, y, x*y)

	//ch := make(chan int, 2) //带缓冲区
	//ch <- 1
	//ch <- 2
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//ch <- 3 //取走之后可继续存放
	//fmt.Println(<-ch)

	//ch := make(chan int)
	//go producer(ch)
	//for v := range ch {
	//	fmt.Println(v)
	//}
	//fmt.Println(<-ch) //关闭的通道打印0值

	//c, quit := make(chan int), make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//fb(c, quit)

	//var wg sync.WaitGroup
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go worker("worker"+strconv.Itoa(i), &wg)
	//}
	//wg.Wait() //等待WaitGroup为空
	//fmt.Println("worker is all done")

	a := 2
	go see(a)
	select { //超时处理
	case <-done:
		fmt.Println("done")
	case <-time.After(time.Second * 1): //监听超时
		fmt.Println("timeout")
	}
}
