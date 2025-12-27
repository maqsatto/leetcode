package leet35

func SearchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	res := 0
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
			res = l
		} else {
			r = mid - 1
		}
	}
	return res
}
