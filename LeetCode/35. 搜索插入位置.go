package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 3, 5, 6}
	var target int
	target = 2
	low := 0
	high := len(nums) - 1
	mid := (high-low)/2 + low
	for low <= high {
		mid = (high-low)/2 + low
		if nums[mid] == target {
			fmt.Println(mid)
			break
			//	return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		}
	}
	if nums[mid] < target {
		return mid + 1
	} else {
		return mid
	}
}
