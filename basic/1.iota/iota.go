package main

import "fmt"

////1.

//func main() {
//	const (
//		a = iota
//		b
//		c
//	)
//	fmt.Printf("%d  %d  %d", a, b, c)
//}
////0 ,1,2

/*2.
func main() {
	const (
		a = 1 << iota
		b = 1 << iota
		c
		d = 1 << iota
		e
	)
	fmt.Printf("%d,%d,%d,%d,%d", a, b, c, d, e)
}
*/
////                    1,2,4,8,16

//如果默认只写变量，执行语句与上方一样

//3
func main() {
	const (
		a = 1         //iota=0
		b             //iota=1
		c             //=2
		d = 1 << iota //=3
		e             //4
		f = iota      //5
		g             //6
	)
	fmt.Printf("%d,%d,%d,%d,%d,%d,%d", a, b, c, d, e, f, g)
}

//                1,1,1,8,16,5,6
