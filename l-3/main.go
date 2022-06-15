package main

import (
	"fmt"
)

func outer() {
	var test6 string = "test"
	fmt.Println(test6)
}

func main() {
	// var 変数名　型　= 値
	var i int = 100
	fmt.Println(i)

	var s string = "aaaaa"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2    string = "aaaa"
		test2 int    = 1999
	)
	fmt.Println(i2, test2)

	var test3 int    // 0
	var test4 string //空文字
	fmt.Println(test3, test4)

	test4 = "aaa"
	fmt.Println(test3, test4)

	// 変数名 := 値

	outer()
}
