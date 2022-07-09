package main

import (
	"fmt"
	"time"
)

// channel
// close

func reciever(name string, ch1 chan int) {
	for {
		i, ok := <-ch1
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}
func main() {
	ch1 := make(chan int, 2)
	/*
		close(ch)

		// fmt.Println(<-ch)

		i, ok := <-ch
		fmt.Println(i, ok)
	*/

	go reciever("1.goroutin", ch1)
	go reciever("2.goroutin", ch1)
	go reciever("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Millisecond)

}
