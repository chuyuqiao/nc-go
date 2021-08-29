package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindGreatestSumOfSubArray([]int{1, -2, 3, 10, -4, 7, 2, -5}))
}

func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	n := len(array)
	dp := make([]int, n)
	dp[0] = array[0]
	maxSum := dp[0]

	for i := 1; i < n; i++ {
		dp[i] = array[i] + max(dp[i-1], 0)
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
