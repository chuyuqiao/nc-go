package main

import "fmt"

func main() {
	reverseKGroup := reverseKGroup(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}, 3)
	for ; reverseKGroup != nil; reverseKGroup = reverseKGroup.Next {
		fmt.Println(reverseKGroup.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @param k int整型
 * @return ListNode类
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}

	return hair.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	node := head
	var pre *ListNode
	for pre != tail {
		tmp := node.Next
		node.Next = pre
		pre = node
		node = tmp
	}
	return tail, head
}
