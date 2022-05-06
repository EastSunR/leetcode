package main

import "fmt"

func main() {

	number1 := N1{num1: 1, num2: 3}
	fmt.Println(number1.jia(number1.num1, number1.num2))
}

type number interface {
	jia(x, y int) int
	jian(x, y int) int
}
type N1 struct {
	num1 int
	num2 int
}

//type N2 struct {
//	num1 int
//	num2 int
//}

func (n1 N1) jia(x, y int) int {
	return n1.num1 + n1.num2
}
func (n1 N1) jian(x, y int) int {
	return n1.num1 - n1.num2
}
