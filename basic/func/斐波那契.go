package main

import "fmt"

func fi(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fi(n-1) + fi(n-2)
	}
}
func main() {
	fmt.Println("请输入一个数")
	n := 0
	fmt.Scanf("%d", &n)
	s := fi(n)
	fmt.Println(s)
}
