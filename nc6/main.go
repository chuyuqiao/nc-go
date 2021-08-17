package main

import (
	"fmt"
	"math"

	"leetcode-go/lib"
)

func main() {
	// write code here
	root := &lib.TreeNode{
		Val: -1,
		Left: &lib.TreeNode{
			Val: 1,
			Left: &lib.TreeNode{
				Val: 2,
			},
			Right: &lib.TreeNode{
				Val: -3,
			},
		},
	}

	maxSum := math.MinInt32
	var maxGain func(*lib.TreeNode) int
	maxGain = func(node *lib.TreeNode) int {
		if node == nil {
			return 0
		}

		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		newPath := node.Val + leftGain + rightGain
		maxSum = max(maxSum, newPath)

		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)

	fmt.Println(maxSum)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
