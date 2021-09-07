package main

import "fmt"

func main() {
	// 123 + 1991 = 2114
	sum := addInList(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		})
	for ; sum != nil; sum = sum.Next {
		fmt.Println(sum.Val)
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head1 ListNode类
 * @param head2 ListNode类
 * @return ListNode类
 */
func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	// 1. 反转链表
	head1 = reverse(head1)
	head2 = reverse(head2)
	// 2. 加法计算
	plus := 0
	node := &ListNode{}
	dummy := node
	for ; head1 != nil && head2 != nil; head1, head2 = head1.Next, head2.Next {
		sum := head1.Val + head2.Val + plus
		plus = sum / 10
		node.Next = &ListNode{Val: sum % 10}
		node = node.Next
	}

	for ; head1 != nil; head1 = head1.Next {
		sum := head1.Val + plus
		plus = sum / 10
		node.Next = &ListNode{Val: sum % 10}
		node = node.Next
	}

	for ; head2 != nil; head2 = head2.Next {
		sum := head2.Val + plus
		plus = sum / 10
		node.Next = &ListNode{Val: sum % 10}
		node = node.Next
	}

	if plus > 0 {
		node.Next = &ListNode{
			Val: plus,
		}
	}

	// 3. 反转链表
	return reverse(dummy.Next)
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	node := head
	for node != nil {
		tmp := node.Next
		node.Next = pre
		pre = node
		node = tmp
	}
	return pre
}
