package main

import "fmt"

func main() {
	//定义
	var m1 map[string]int
	fmt.Println(m1 == nil)
	if m1 == nil {
		m1 = make(map[string]int)
	}
	m1["zhu"] = 520
	m1["hong"] = 1314
	fmt.Println(m1)
	fmt.Println("---------------------------------")
	var m3 = make(map[string]int) //m3 := make(map[string]int)
	fmt.Println(m3 == nil)
	fmt.Println("---------------------------------")
	m2 := map[string]int{"zhao": 1, "qian": 2, "sun": 3}
	fmt.Println(m2 == nil)
	fmt.Println(m2)
	fmt.Println(m2["zhao"])
	delete(m2, "zhao")
	fmt.Println(m2)
	i, ok := m2["zhao"]
	if ok {
		fmt.Println("值：", i)
	} else {
		fmt.Println("不存在")
	}
	i, ok = m2["qian"]
	if ok {
		fmt.Println("值：", i)
	} else {
		fmt.Println("不存在")
	}
}
