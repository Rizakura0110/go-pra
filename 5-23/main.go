package main

import "fmt"

// グローバルに宣言
const Pi = 3.14

const (
	URL      = "http://example.xxx.xxx"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota // 連続する整数
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	/* error
	Pi = 3
	fmt.Println(Pi)
	*/

	fmt.Println(URL, SiteName)

	// 1 1 1 A A A となる
	fmt.Println(A, B, C, D, E, F)

	fmt.Println(c0, c1, c2)
}
