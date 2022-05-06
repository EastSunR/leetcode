package main

import "fmt"

func main() {
	zhao := Husband{name: "赵东昕", wife: "朱红瑜", year: 20}
	zhu := Wife{name: "朱红瑜", husband: "赵东昕", year: 18}
	he(zhao)
	fmt.Println()
	he(zhu)
}

type USB interface {
	st()
	end()
}
type Husband struct {
	name string
	wife string
	year int
}
type Wife struct {
	name    string
	husband string
	year    int
}

func (zhao Husband) st() {
	fmt.Println(zhao.name, "的年龄是", zhao.year)
}
func (zhao Husband) end() {
	fmt.Println(zhao.name, "爱老婆", zhao.wife)
}
func (zhu Wife) st() {
	fmt.Println("漂亮小仙女", zhu.name, "的年龄是", zhu.year)
}
func (zhu Wife) end() {
	fmt.Println(zhu.name, "爱老公", zhu.husband)
}
func he(mua USB) {
	mua.st()
	mua.end()
}
