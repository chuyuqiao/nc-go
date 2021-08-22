package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(isSymmetric(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}))
}

func isSymmetric(root *TreeNode) bool {
	// write code here
	return isSymmetricBoth(root, root)
}

func isSymmetricBoth(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricBoth(left.Left, right.Right) && isSymmetricBoth(left.Right, right.Left)
}
