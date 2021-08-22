package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// write code here
	root := &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: -3,
			},
		},
	}

	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
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
