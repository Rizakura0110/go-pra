package main

import "fmt"

// クロージャー
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}
func main() {
	f := Later()
	fmt.Println(f("test"))
}
