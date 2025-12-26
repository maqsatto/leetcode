package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	var check int = 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[check-2] {
			nums[check] = nums[i]
			check++
		}
	}
	fmt.Println(nums)
}
