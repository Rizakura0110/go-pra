package main

import "fmt"

func main(){
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	// float64
	fl:=3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n" , fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n" , fl32)

	fmt.Printf("%T\n" , float64(fl32))

	var t, f bool = true, false
	fmt.Println(t, f)

	var s string = "hello"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(
		`test
				test
					 test`)
	fmt.Println("\"")
	fmt.Println('"')
	fmt.Println(s[0]) //byte型
	fmt.Println(string(s[0])) //文字列
}