package main

import "fmt"

func main() {
	fmt.Println(subsequence(1, []int{-1}))
	fmt.Println(subsequence(3, []int{1, 2, 3}))
	fmt.Println(subsequence(4, []int{4, 2, 3, 5}))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算
 * @param n int整型 数组的长度
 * @param array int整型一维数组 长度为n的数组
 * @return long长整型
 */
func subsequence(n int, array []int) int64 {
	// write code here
	dp := make([]int64, n+1)
	dp[0] = 0
	dp[1] = max(int64(array[0]), 0)

	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-2]+int64(array[i-1]), dp[i-1])
	}

	return dp[n]
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
