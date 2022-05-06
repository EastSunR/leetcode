package main

import (
	"fmt"
	"sort"
)

func main() {
	m1 := map[int]string{10: "zhao", 30: "sun", 20: "qian", 40: "li", 50: "zhou"}
	fmt.Println(m1)
	for k, v := range m1 {
		println(k, "---->", v)
	}
	fmt.Println("----------------------------------")
	key := make([]int, 0, len(m1))
	for i, _ := range m1 {
		key = append(key, i)
	}
	for _, i := range key {
		fmt.Println(i, "--->", m1[i])
	}
	fmt.Println("---------")
	sort.Ints(key)
	for _, i := range key {
		fmt.Println(i, "--->", m1[i])
	}
}
