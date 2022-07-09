package main

import "fmt"

// channel
// 宣言操作
func main() {
	// 双方向
	var chl chan int // nil
	chl = make(chan int)
	ch2 := make(chan int)
	fmt.Println(cap(chl))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)

	ch3 <- 1
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))

	i := <-ch3
	fmt.Println(i)
	fmt.Println("len", len(ch3))
	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))

	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))
	// 受信専用
	//var ch2 <-chan int

	// 送信専用
	//var ch3 chan<- int

	/* error
	ch3 <- 2
	ch3 <- 3
	ch3 <- 2
	ch3 <- 3
	ch3 <- 2
	ch3 <- 5
	*/
}
