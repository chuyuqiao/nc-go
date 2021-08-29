package main

import "fmt"

func main() {
	fmt.Println(knapsack(
		10,
		2,
		[][]int{
			{
				1, 3,
			},
			{
				10, 4,
			},
		},
	))
}

func knapsack(V int, n int, vw [][]int) int {

	/*
		   0  1  2  3  4  5  6  7  8  9  10
		0  0  3 -1 -1 -1 -1 -1 -1 -1 -1 -1
		1  0  3 -1 -1 -1 -1 -1 -1 -1 -1 4
	*/

	// write code here
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, V+1)
		for j := 0; j < V+1; j++ {
			dp[i][j] = -1
		}
	}

	dp[0][0] = 0
	if vw[0][0] < V {
		dp[0][vw[0][0]] = vw[0][1]
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= V; j++ {
			if dp[i-1][j] >= 0 {
				dp[i][j] = dp[i-1][j]
			}
		}
		for j := 0; j <= V-vw[i][0]; j++ {
			if dp[i-1][j] >= 0 {
				w := dp[i-1][j] + vw[i][1]
				if w > dp[i][j+vw[i][0]] {
					dp[i][j+vw[i][0]] = w
				}
			}
		}
	}

	maxValue := -1
	for i := 0; i <= V; i++ {
		if dp[n-1][i] > maxValue {
			maxValue = dp[n-1][i]
		}
	}

	return maxValue
}
