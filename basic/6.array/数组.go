package main

import "fmt"

func main() {
	//1
	var a = [3]int{1, 2, 3}
	fmt.Println(a)
	//2
	var b = [5]int{2: 5, 4: 6}
	fmt.Println(b)
	fmt.Println(len(b))
	//3
	var c = [...]int{2, 4, 6, 8}
	fmt.Println(c)
	fmt.Println(len(c)) //长度
	fmt.Println(cap(c)) //容量
	d := [4]int{1, 3, 4, 5}
	fmt.Println(d)
	//for range
	for x, y := range d {
		fmt.Printf("下标是：%d  值为：%d\n", x, y)
	}
	e := [3][5]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(e)
	for x, y := range e {
		for z, num := range y {
			fmt.Printf("%d %d %d  **", x, z, num)
		}
		fmt.Println()
	}
}
