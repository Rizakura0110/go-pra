package main

import "fmt"

// copy
func main() {
	sl := []int{100, 220}
	sl2 := sl
	sl2[0] = 1000
	fmt.Println(sl) // 1000 220

	sl3 := []int{100,200}
	sl4 := make([]int, 5,10)
	n := copy(sl4, sl3)
	fmt.Println(n, sl4)

}
