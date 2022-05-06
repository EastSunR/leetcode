package main

import "sort"

func main() {
	var nums []int
	for i, v := range nums {
		nums[i] = v * v
	}
	sort.Ints(nums)
	return nums
}
