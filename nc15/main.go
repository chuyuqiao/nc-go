package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(levelOrder(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
}

func levelOrder(root *TreeNode) [][]int {
	// write code here
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	var res [][]int
	for len(q) > 0 {
		n := len(q)
		var level []int
		for ; n > 0; n-- {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
