package leet26

func RemoveDuplicates(nums []int) int {
	var check int = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[check] = nums[i]
			check++
		}
	}
	return len(nums)
}
