package main

import "fmt"

// ジェネレーター

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	ints := integers()

	// 1 2 3 4 5
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	// 1 2 3 4 5
	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
}
