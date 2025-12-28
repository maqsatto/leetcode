package leet69

import "math"

func MySqrt(x int) int {
	if x <= 0 {
		return 0
	}
	z := float64(x)
	for i := 0; i < 20; i++ {
		prev := z
		z = z - (z*z-float64(x))/(2*z)
		if math.Abs(z-prev) < 1e-9 {
			break
		}
	}
	return int(z)
}
