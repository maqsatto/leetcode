package main

import (
	"fmt"
	"math/rand"
	"time"

	leet74 "github.com/maqsatto/leetcode/golang/medium/leet74_Search_a_2D_Matrix"
)

func generateMatrix(m, n int) [][]int {
	matrix := make([][]int, m)
	val := 1

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = val
			val++
		}
	}
	return matrix
}
func generateRandomMatrix(m, n int) [][]int {
	rand.Seed(time.Now().UnixNano())

	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = rand.Intn(m * n * 10)
		}
	}
	return matrix
}

func main() {
	matrix := generateRandomMatrix(100000, 10000)
	start1 := time.Now()
	leet74.SearchMatrix(matrix, 100000)
	fmt.Println(time.Since(start1).Seconds())

}
