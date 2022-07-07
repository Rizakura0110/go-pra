package main

import "fmt"

// for
// 可変長引数

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
func main() {
	sl := []string{"apple", "banana"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(Sum())

	sl3 := []int{1, 2, 3}
	fmt.Println(Sum(sl3...))
}
