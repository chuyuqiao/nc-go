package main

import "fmt"

func main() {
	fmt.Println(isPail(array2List([]int{1})))
	fmt.Println(isPail(array2List([]int{1, 2, 2})))
	fmt.Println(isPail(array2List([]int{1, 2, 2, 1})))
	fmt.Println(isPail(array2List([]int{323809, -861952, 178712, 13683, 162169, 620654, 483476, 882177, -797062, -499208, 216105, -59351, -307317, -25191, 61574, 645960, 511402, 814708, 172894, 125933, -584712, -559044, -924916, 582034, -295024, 495888})))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类 the head
 * @return bool布尔型
 */
func isPail(head *ListNode) bool {
	// write code here
	slow, fast := head, head
	for ; fast.Next != nil && fast.Next.Next != nil; fast = fast.Next.Next {
		slow = slow.Next
	}
	pHead := slow.Next
	slow.Next = nil
	pHeadReversed := reverse(pHead)
	for ; head != nil && pHeadReversed != nil; head, pHeadReversed = head.Next, pHeadReversed.Next {
		if head.Val != pHeadReversed.Val {
			return false
		}
	}
	if head == pHeadReversed {
		return true
	}
	if head != nil && head.Next == nil {
		return true
	}
	return false
}

func reverse(head *ListNode) *ListNode {
	node := head
	var pre *ListNode
	for node != nil {
		tmp := node.Next
		node.Next = pre
		pre = node
		node = tmp
	}
	return pre
}

func array2List(arr []int) *ListNode {
	fmt.Println(len(arr))
	node := &ListNode{Val: arr[0]}
	dummy := &ListNode{Next: node}
	for i := 1; i < len(arr); i++ {
		if i == len(arr)/2 {
			fmt.Println(arr[i])
		}
		node.Next = &ListNode{Val: arr[i]}
		node = node.Next
	}
	return dummy.Next
}
