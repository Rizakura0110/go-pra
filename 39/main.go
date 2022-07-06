package main

import "fmt"

// switch
// 型スイッチ
// 型アサーション
// interface型
func anything(a interface{}) {
	fmt.Println(a)
}
func main() {
	anything("aaa")
	anything(1)

	// 3 true
	var x interface{} = 3
	i, isInt := x.(int)
	fmt.Println(i, isInt) //5
	// error fmt.Println(x + 3)

	// 0 false
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("none")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I dont know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("none")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "is bool")
	case int:
		fmt.Println(v, "is int")
	default:
	}
}
