package main

import "fmt"
func main() {
	fmt.Println("请输入两个数")
	var c, b int
	fmt.Scanf("%d %d", &c, &b)
	fmt.Println(add(c, b))
}
func add(a, b int) (int, int) {
	an := func(x, y int) int {
		return x + y
	}
	//	fmt.Println(sll)
	return an(a, b), a + b

}
