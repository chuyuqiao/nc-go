package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(sortedArrayToBST([]int{-1, 0, 1, 2}))
}

func sortedArrayToBST(num []int) *TreeNode {
	// write code here
	if len(num) == 0 {
		return nil
	}
	return dfs(num, 0, len(num)-1)
}

func dfs(num []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end + 1) / 2
	root := &TreeNode{Val: num[mid]}
	root.Left = dfs(num, start, mid-1)
	root.Right = dfs(num, mid+1, end)
	return root
}
