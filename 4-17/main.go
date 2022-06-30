package main

import "fmt"

// 配列
func main(){
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"a", "b"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2 ,3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "B"} // 中身の数に要素数が引っ張られる
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[1])

	arr2[2] = "x"
	fmt.Println(arr2[2])

	/* error
	var arr5 [4]int
	arr5 = arr1
	fmt.Println(arr5)
	*/

	var arr5 [3]int
	arr5 = arr1
	fmt.Println(arr5)

	fmt.Println(len(arr5))
}