package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	new := make([]int, len(nums))
	for i, _ := range nums {
		new[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, new)
	fmt.Println(nums)
}
