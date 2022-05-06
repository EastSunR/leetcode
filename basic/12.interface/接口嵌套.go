package main

import "fmt"

func main() {
	var m1 C
	m1 = n1{}
	m1.t1()
	m1.t2()
	m1.t3()
}

type A interface {
	t1()
}
type B interface {
	t2()
}
type C interface {
	A
	B
	t3()
}
type n1 struct {
}

func (n n1) t1() {
	fmt.Println("111111")
}
func (n n1) t2() {
	fmt.Println(222)
}
func (n n1) t3() {
	fmt.Println(333)
}
