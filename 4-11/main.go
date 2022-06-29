package main

import "fmt"

func main(){
	var i int = 100

	fmt.Println(i + 50)

	fmt.Printf("%T\n", i)

	fmt.Printf("%T\n", int32(i))
}