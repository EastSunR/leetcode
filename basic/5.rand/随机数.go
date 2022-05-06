package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a int
	//两种改变随机数大小的方法
	a = rand.Int() % 20
	fmt.Println(a)
	a = rand.Intn(10)
	fmt.Println(a)
	//时间函数
	t := time.Now()
	fmt.Println(t)
	//将时间转化为整数 到1970年一月一日00:00:00 的时间（秒/纳秒）
	t2 := t.Unix()     //秒
	t3 := t.UnixNano() //纳秒
	fmt.Println(t2)
	fmt.Println(t3)
	//生成种子函数
	rand.Seed(10)
	n := rand.Intn(10) + 10 //10-20
	fmt.Println(n)          //随机数随种子改变，由于种子固定 ，生成随机数固定	14
	//采用时间作为种子来改变
	fmt.Println("***")
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < 11; i++ {
		num := rand.Intn(90) + 10
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
