package main

import (
	"fmt"

	"github.com/maqsatto/leetcode/golang/medium/leet189"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	leet189.Rotate(nums, 2)
	fmt.Println(nums)
}
