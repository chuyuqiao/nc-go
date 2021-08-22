package main

import "fmt"

func main() {
	fmt.Println(getLongestPalindrome("baabccc", 7))
}

func getLongestPalindrome(A string, n int) int {
	// write code here
	maxLen := 0
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			xLen := j - i + 1
			if A[i] == A[j] {
				if xLen <= 2 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == 1 {
				maxLen = max(maxLen, xLen)
			}
		}
	}
	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
