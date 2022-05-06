package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "hello world"
	//1.Strings.Contains
	//是否包含指定内容
	f := strings.Contains(s, "llo")
	fmt.Println(f)
	f = strings.Contains(s, "nihao")
	fmt.Println(f)
	fmt.Println("---------------------------------------------------")
	//2.strings.ContainsAny()
	//判断所给指定字符串中是否存在字符在原字符串中（s）出现
	f = strings.ContainsAny(s, "abc")
	fmt.Println(f)
	f = strings.ContainsAny(s, "abcd")
	fmt.Println(f)
	fmt.Println("-----------------------------------------")
	//3.strings.Count()
	//统计所给字符串（字符）出现的次数
	//s  = "hello world"
	var num int
	num = strings.Count(s, "l")
	fmt.Println(num)
	num = strings.Count(s, "llo")
	fmt.Println(num)
	fmt.Println("-----------------------------")
	//4.
	s1 := "20220429.txt"
	f = strings.HasPrefix(s1, "2022")
	fmt.Println(f)
	f = strings.HasSuffix(s1, "txt")
	fmt.Println(f)
	fmt.Println("------------------------------")
	//5.strings.Index()
	//"hello world"
	fmt.Println(strings.Index(s, "l"))
	fmt.Println(strings.Index(s, "le"))
	fmt.Println("strings.IndexAny():")
	fmt.Println(strings.IndexAny(s, "abc"))
	fmt.Println(strings.IndexAny(s, "sbcd"))
	fmt.Println(strings.IndexAny(s, "abcde"))
}
