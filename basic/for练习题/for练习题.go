package main

import "math"

///一。
//func main() {
//	var i int
//	for i = 58; i >= 23; i-- {
//		fmt.Print(i, " ")
//	}
//	fmt.Println("*******************")
//	sum := 0
//	for i = 1; i <= 100; i++ {
//		sum += i
//	}
//	fmt.Println("sum= ", sum)
//	fmt.Println("*******************")
//	cnt := 0
//	for i = 1; i < 100; i++ {
//		if i%3 == 0 && i%5 != 0 {
//			cnt++
//			fmt.Printf("%d ", i)
//			if cnt%5 == 0 && cnt != 0 {
//				fmt.Println()
//			}
//		}
//
//	}
//}
/*****************************/
//二.九九乘法表
/*
	func main() {
		for i := 1; i <= 9; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("%d*%d=%d ", i, j, i*j)
			}
			fmt.Println()
		}
	}
*/
//*************************
func main() {

	for i := 100; i < 1000; i++ {
		ge := i % 10
		shi := i / 10 % 10
		bai := i / 100
		if math.Abs((math.Pow(float64(ge), 3)+math.Pow(float64(shi), 3)+math.Pow(float64(bai), 3))-float64(i)) < 1e-8 {
			println(i)
		}
	}
}
