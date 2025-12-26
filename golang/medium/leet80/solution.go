package leet80

import "fmt"

func RemoveDuplicates(nums []int) int {

	var check int = 2
	for i := range nums {
		if nums[i] != nums[check-2] {
			nums[check] = nums[i]
			check++
		}
	}
	fmt.Println(nums)
	return check
}
