package leet27

import "slices"

func RemoveElement(nums []int, val int) int {
	nums = slices.DeleteFunc(nums, func(e int) bool {
		return e == val
	})
	return len(nums)
}
