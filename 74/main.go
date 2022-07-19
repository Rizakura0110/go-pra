package main

import (
	"74/foo"
	"fmt"
	. "fmt"
	f "fmt"
)

// スコープ

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.8"
	return AppName + "" + Version
}

func Do(s string) (b string) {
	//var b string = s
	b = s
	{
		b := "bbbb"
		fmt.Println(b)
	}
	Println(b)
	return b
}
func main() {
	f.Println(foo.Max)
	//fmt.Println(foo.min) error

	Println(foo.ReturnMin())

	fmt.Println(appName())
	//fmt.Println(AppName, Verion)
	fmt.Println(Do("aaaa"))
}
