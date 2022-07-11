package main

import "fmt"

type T struct {
	User
}

type User struct {
	Name string
	Age  int
}

func (u *User) setName() {
	u.Name = "a"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	fmt.Println(t.Name)

	t.User.setName()
	fmt.Println(t)
}
