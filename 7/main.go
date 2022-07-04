package main

import "fmt"

//func Plus(x int, y int) int {
func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return result
}

func Noreturn() {
	fmt.Println("No Return")
}

// 関数を引数に取る関数
func CallFunction(f func()) {
	f()
}

func main() {
	i := Plus(3, 4)
	fmt.Println(i)

	i2, i3 := Div(9, 4)
	fmt.Println(i2, i3)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()

	CallFunction(func() {
		fmt.Println("test")
	})

}
