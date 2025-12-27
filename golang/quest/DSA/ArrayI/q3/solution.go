package q3

func FindMaxConsecutiveOnes(nums []int) int {
	var counter int32
	var max int32
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			counter++
			if counter > max {
				max = counter
			}
		} else {
			counter = 0
		}
	}
	return int(max)
}
