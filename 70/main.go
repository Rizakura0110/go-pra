package main

import "fmt"

// interface

type Stringfy interface {
	ToString() string
}
type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "taro", Age: 21},
		&Car{Number: "1234565", Model: "ab-123"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}
