package main

import (
	"fmt"
)

func main() {
	var nums = []int{-1, 0, 3, 5, 9, 12}
	var target int
	fmt.Scanf("%d", &target)
	/*
		//1.普通暴力

		n := -1
		for x, v := range nums {
			if target == v {
				n = x
			}
		}
	*/
	//2.
	low, high := 0, len(nums)-1
	for low < high {
		mid := (high-low)/2 + low
		if nums[mid] == target {
			fmt.Println(mid)
			break
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println(-1)
}
