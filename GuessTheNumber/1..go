package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100) + 1
	fmt.Println("请输入一个1-100的整数")
	var t int
	fmt.Scan(&t)
	count := 1
	min, max := 0, 100
	for t != n {
		if t > n {
			fmt.Println(t, "太大了")
			fmt.Println("提示范围：", min, "-", t)
			max = t
		} else {
			fmt.Println(t, "太小了")
			fmt.Println("提示范围：", t, "-", max)
			min = t
		}
		count++
		fmt.Printf("\n")
		fmt.Println("请再次输入：")
		fmt.Scan(&t)
		for t <= min || t >= max {
			fmt.Println("输入错误,请重新输入")
			fmt.Println("正确范围为", min, "-", max)
			fmt.Scan(&t)
		}
	}
	fmt.Println("恭喜您正确猜对，共用了", count, "次猜出")
}
