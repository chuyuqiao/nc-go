package main

import "fmt"

func main() {
	duplicates := deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	})
	for duplicates != nil {
		fmt.Println(duplicates.Val)
		duplicates = duplicates.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @return ListNode类
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return nil
	}
	if head.Next != nil && head.Next.Val == head.Val {
		for head.Next != nil && head.Next.Val == head.Val {
			head = head.Next
		}
		return deleteDuplicates(head.Next)
	}
	head.Next = deleteDuplicates(head.Next)
	return head
}
