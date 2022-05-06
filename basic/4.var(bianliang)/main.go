package main

import "fmt"

func main() {
	//1.
	var a int = 1
	var s string = "我爱你"
	fmt.Println(a, s)
	//2.
	var b, c, d = 1.1, 2, "hello!"
	fmt.Printf("b=%.2f c=%d d=%s\n", b, c, d)
	//3.
	e := 3
	f := "world"
	fmt.Println(e, f)
	//4
	var (
		g = 5
		h = 5.5
		i = "hi"
	)
	fmt.Println(g, h, i)
}
