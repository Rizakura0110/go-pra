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

	// for
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	}

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

	arr2 := [3]int{1, 2, 3}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}

	arr3 := [3]int{1, 2, 3}
	for _, v := range arr3 {
		fmt.Println(v)
	}

	sl := []string{"test", "pc", "apple"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}

	for k := range m {
		fmt.Println(k)
	}

	// switch
	/*
		n := 8
		switch n {
		case 1, 2:
			fmt.Println("1 or 2")
		case 3, 4:
			fmt.Println("3 or 4")
		default:
			fmt.Println("no")

		}
	*/

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("no")
	}
}
