package main

import "fmt"

func main() {
	fmt.Println(LCS("1A2C3D4B56", "B1D23A456A"))
}

/**
 * longest common subsequence
 * @param s1 string字符串 the string
 * @param s2 string字符串 the string
 * @return string字符串
 */
func LCS(s1 string, s2 string) string {
	// write code here
	m, n := len(s1), len(s2)
	if m == 0 || n == 0 {
		return "-1"
	}
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	b := make([][]int, m+1)
	for i := range b {
		b[i] = make([]int, n+1)
	}

	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
				b[i+1][j+1] = 1 //上方
			} else {
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
					b[i+1][j+1] = 2 //左方
				} else {
					dp[i+1][j+1] = dp[i+1][j]
					b[i+1][j+1] = 3 //上方
				}
			}
		}
	}
	res := ans(s1, s2, m, n, b)
	if len(res) == 0 {
		return "-1"
	}
	return res
}

func ans(s1, s2 string, m, n int, b [][]int) (res string) {
	if m == 0 || n == 0 {
		return
	}
	if b[m][n] == 1 {
		res += ans(s1, s2, m-1, n-1, b)
		res += s1[m-1 : m]
	} else if b[m][n] == 2 {
		res += ans(s1, s2, m-1, n, b)
	} else if b[m][n] == 3 {
		res += ans(s1, s2, m, n-1, b)
	}
	return
}
