//先进后出
package main

import (
	"fmt"
	"os"
)

func main() {
	//defer func() {
	//	fmt.Println("第一")
	//}()
	//defer func() {
	//	fmt.Println("第二")
	//}()
	//fmt.Println("hello world")
	defer func() {
		fmt.Println(111)
	}()
	fmt.Println(555)
	os.Exit(1)
}
