package leet88

import (
	"fmt"
	"sort"
)

func Merge(nums1 []int, m int, nums2 []int, n int) {
	newArray := []int{}
	for i := 0; i < m; i++ {
		newArray = append(newArray, nums1[i])
	}
	for i := 0; i < n; i++ {
		newArray = append(newArray, nums2[i])
	}

	sort.Ints(newArray)
	for i := 0; i < n+m; i++ {
		nums1[i] = newArray[i]
	}
	fmt.Println(nums1)
}
