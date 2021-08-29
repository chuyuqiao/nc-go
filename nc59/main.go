package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{
		{
			1, 3, 5, 9,
		},
		{
			8, 1, 3, 4,
		},
		{
			5, 0, 6, 1,
		},
		{
			8, 8, 4, 0,
		},
	}))
}

func minPathSum(matrix [][]int) int {
	// write code here
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	msum := 0
	for i := 0; i < m; i++ {
		msum += matrix[i][0]
		dp[i][0] = msum
	}
	nsum := 0
	for i := 0; i < n; i++ {
		nsum += matrix[0][i]
		dp[0][i] = nsum
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = matrix[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
