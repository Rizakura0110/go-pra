package main

import (
	"fmt"
	"os"
)

// defer
func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("start")
}

func RunDefer() {
	// 後から登録された文から処理
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}
func main() {
	TestDefer()

	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))
}
