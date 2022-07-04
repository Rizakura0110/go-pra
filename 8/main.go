package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("test")
	}

	if b := 100; b == 100 {
		fmt.Println("hundred")
	}

	var s string = "a"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)
}
