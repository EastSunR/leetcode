package main

import "fmt"

func sum(num ...int) int {
	var s int = 0
	for i := 0; i < len(num); i++ {
		s += num[i]
	}
	return s
}
func main() {
	fmt.Println(sum(1, 3, 5, 7, 9))
}
