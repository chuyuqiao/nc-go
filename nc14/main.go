package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(Print(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}))
}

func Print(pRoot *TreeNode) [][]int {
	// write code here
	if pRoot == nil {
		return nil
	}
	q := []*TreeNode{pRoot}
	var res [][]int

	for len(q) > 0 {
		n := len(q)
		var level []int
		for ; n > 0; n-- {
			node := q[0]
			p := q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
			q = p
		}
		res = append(res, level)
	}

	n := len(res)
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			in := len(res[i])
			for j := 0; j < in/2; j++ {
				res[i][j], res[i][in-j-1] = res[i][in-j-1], res[i][j]
			}
		}
	}
	return res
}
