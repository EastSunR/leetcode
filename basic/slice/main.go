package main

import "fmt"

func main() {
	//普通定义
	var s1 = []int{1, 2, 3}
	fmt.Printf("%T\n", s1)
	fmt.Println(s1)
	s2 := []int{2, 4, 6, 8}
	fmt.Println(len(s2))
	fmt.Println(s2)
	//make定义
	s3 := make([]int, 3, 4)
	fmt.Println("切片长度为：", len(s3))
	fmt.Println("切片容量为：", cap(s3))
	for i := 1; i <= 3; i++ {
		s3[i-1] = i
	}
	//append追加切片元素
	s3 = append(s3, 4, 5)
	fmt.Println(s3)
	fmt.Println("切片长度变为：", len(s3))
	fmt.Println("切片容量变为：", cap(s3))
	fmt.Println("**************************")
	//append ,copy为深拷贝（只拷贝数组，改变任意一个，另一个不会变化）
	s3 = append(s3, s2...)
	fmt.Println("s3为：", s3)
	fmt.Println("s2:", s2)
	fmt.Println("s3切片长度变为：", len(s3))
	fmt.Println("s3切片容量变为：", cap(s3))
	fmt.Printf("****************************\n")
	s2[0] = 1000
	fmt.Println("s3：", s3)
	fmt.Println("s2:", s2)
	copy(s3, s2)
	fmt.Println("s3：", s3)
	fmt.Println("s2:", s2)
	s2[0] = 123456
	fmt.Println("-------------------")
	fmt.Println("s3：", s3)
	fmt.Println("s2:", s2)
	fmt.Println("-----------")
	//若直接相等，切片拷贝的是数组地址为浅拷贝
	s4 := s2
	fmt.Println("s2:", s2)
	fmt.Println("s4:", s4)
	s2[0] = 1
	fmt.Println("s2:", s2)
	fmt.Println("s4:", s4)
}
