package main

import (
	"fmt"
)

func main() {
	type Person struct {
		name string
		age  int
	}
	type Student struct {
		*Person
		Number int
	}
	p := &Person{
		name: "zhuhongyu",
		age:  19,
	}
	s := Student{
		Person: p,
		Number: 5,
	}
	fmt.Printf("%T\n", p)
	fmt.Println(p)
	fmt.Println("-----------------")
	fmt.Println(s)
	println("你好")
}
