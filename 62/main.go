package main

import "fmt"

// 構造体
type User struct {
	Name string
	Age  int
	// X, Y int
}

func updateUser(user User) {
	user.Name = "a"
	user.Age = 1000
}

func updateUser2(user *User) {
	user.Name = "a"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7)

	user8 := &User{}
	fmt.Println(user8)

	updateUser(user1)
	updateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)
}
