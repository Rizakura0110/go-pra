package main

import "fmt"

func main(){
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = 3.34
	fmt.Println(x)

	x = "hello"
	fmt.Println(x)

	x = [3]string{"apple", "banana", "orange"}
	fmt.Println(x)
	
	/* error
	x = 2
	fmt.Println(x + 3)
	*/
}
//fmt.Printf("%T\n"