package main

import (
	"fmt"
	"sort"
)

func main() {
	m1 := map[string]int{"Acc": 11, "d": 45, "cdf": 456, "cfg": 546, "FSFJ": 4545}
	fmt.Println(m1, "\n", "----------------------")
	key := make([]string, 0, len(m1))
	for i, _ := range m1 {
		key = append(key, i)
	}
	fmt.Println(key)
	sort.Strings(key)
	for _, i := range key {
		println(i, "----->", m1[i])
	}
}
