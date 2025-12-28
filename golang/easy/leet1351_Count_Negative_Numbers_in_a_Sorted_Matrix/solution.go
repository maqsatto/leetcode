package leet1351

import "fmt"

func countNegatives(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	i := 0
	j := n - 1
	count := 0

	for i < m && j >= 0 {
		if grid[i][j] < 0 {
			// все элементы ниже в этом столбце тоже < 0
			count += m - i
			j-- // идем левее
		} else {
			i++ // идем ниже
		}
	}
	return count
}
func TestCountNegatives() {
	// [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
	grid := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	fmt.Println(countNegatives(grid))
}
