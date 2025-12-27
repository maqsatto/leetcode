package q1

import "sort"

func FindErrorNums(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == i {
			continue
		} else if nums[i] < i {
			return []int{nums[i], i + 1}
		} else if nums[i] > i {
			return []int{nums[i], i + 1}
		}
	}
	return []int{}
}
