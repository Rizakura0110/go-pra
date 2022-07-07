package main

import "fmt"

// マップ
func main() {
	var m = map[string]int{"a": 100, "b": 200}

	fmt.Println(m)

	m2 := map[string]int{"a": 100, "b": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "a",
		2: "b",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)
	m4[1] = "test"
	m4[3] = "test3"
	fmt.Println(m4)

	fmt.Println(m["a"])
	fmt.Println(m4[4]) // 空文字

	s, ok := m4[4]
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(s, ok)

	m4[3] = "changetest"
	fmt.Println(m4[3])

	delete(m4, 3)
	fmt.Println(m4[3])

	// map for
	m6 := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m6 {
		fmt.Println(k, v)
	}
}
