package leet189

import "fmt"

func Rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	if k == 0 {
		return
	}
	k = k % len(nums)

	mysl := make([]int, 0, len(nums))
	mysl = append(mysl, nums[len(nums)-k:]...)
	fmt.Println(mysl)
	mysl = append(mysl, nums[:len(nums)-k]...)
	copy(nums, mysl)
}
