package main

import "fmt"

func main() {
	//	num := 6 + 3.5i //默认complex128
	var num complex64 = 6 + 3.5i
	fmt.Printf("%f\n", num) ///complex由两个float组成所以输出使用%f
	var shi float32 = real(num)
	var xu float32 = imag(num) //必须为complex的一半
	fmt.Printf("shi=%f\txu=%f", shi, xu)
}
