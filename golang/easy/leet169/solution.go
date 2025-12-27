package leet169

func MajorityElement(nums []int) int {
	mymap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mymap[nums[i]]++
	}
	max := 0
	maxValue := 0
	for key, value := range mymap {
		if value > maxValue {
			max = key
			maxValue = value
		}
	}
	return max
}
