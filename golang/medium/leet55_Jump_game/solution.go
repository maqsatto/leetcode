package leet55

func CanJump(nums []int) bool {
	far := 0

	for i := 0; i < len(nums); i++ {
		if i > far {
			return false
		}
		if i+nums[i] > far {
			far = i + nums[i]
		}
		if far >= len(nums)-1 {
			return true
		}
	}
	return true
}

//func CanJump(nums []int) bool {
//	if len(nums) < 2 {
//		return true
//	}
//	i := 0
//	for i < len(nums) {
//		if i >= len(nums)-1 {
//			return true
//		}
//		for j := i + 1; j < nums[i]; j++ {
//			if nums[j] == 0 {
//				if j == len(nums)-1 {
//					return true
//				}
//			}
//			if j >= len(nums)-1 {
//				return true
//			}
//			i = j
//		}
//		if nums[i] == 0 {
//			if i == len(nums)-1 {
//				return true
//			}
//			return false
//		}
//		i += nums[i]
//	}
//	return true
//}
