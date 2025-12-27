package q2

func Shuffle(nums []int, n int) []int {
	array := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		array = append(array, nums[i])
		array = append(array, nums[i+n])
	}
	return array
}
