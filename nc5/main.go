package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 0}}
	fmt.Println(sumNumbers(&root))
}

/**
 *
 * @param root TreeNode类
 * @return int整型
 */
func sumNumbers(root *TreeNode) int {
	// write code here
	if root == nil {
		return 0
	}

	sum := 0
	q := []*TreeNode{root}
	r := []int{root.Val}
	for len(q) > 0 {
		node := q[0]
		q1 := q[1:]
		nodeV := r[0]
		r1 := r[1:]
		if node.Left == nil && node.Right == nil {
			sum += nodeV
		} else {
			if node.Left != nil {
				q1 = append(q1, node.Left)
				r1 = append(r1, nodeV*10+node.Left.Val)
			}
			if node.Right != nil {
				q1 = append(q1, node.Right)
				r1 = append(r1, nodeV*10+node.Right.Val)
			}
		}
		q = q1
		r = r1
	}
	return sum
}
