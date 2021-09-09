package main

import "fmt"

func main() {
	merged := Merge(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
			},
		},
	}, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
			},
		},
	})

	for ; merged != nil; merged = merged.Next {
		fmt.Println(merged.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}

	dummy := &ListNode{}
	tail := dummy
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val < pHead2.Val {
			tail.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			tail.Next = pHead2
			pHead2 = pHead2.Next
		}
		tail = tail.Next
	}

	if pHead1 == nil {
		tail.Next = pHead2
	}

	if pHead2 == nil {
		tail.Next = pHead1
	}

	return dummy.Next
}
