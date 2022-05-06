//数组指针  var *p [n]int   存储了一个n大小的数组的地址的指针
//首先是一个指针 其中含有一个地址  是一个数组的地址
package main

import "fmt"

func main() {

	a := [4]int{1, 2, 3, 4}
	fmt.Println(a)
	p := &a
	//var p *[4]int
	//p = &a
	fmt.Println(p)
	(*p)[1] = 9999 //==         p[1]=999
	fmt.Println(a)
	fmt.Println(p[1])
}
