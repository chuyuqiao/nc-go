package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(maxDepth(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
}

func maxDepth(root *TreeNode) int {
	// write code here
	if root == nil {
		return 0
	}
	q, depth := []*TreeNode{root}, 0
	for len(q) > 0 {
		n := len(q)
		for ; n > 0; n-- {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		depth++
	}
	return depth
}
