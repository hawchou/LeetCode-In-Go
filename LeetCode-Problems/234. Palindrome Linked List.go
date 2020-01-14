package main


import (
	. "LeetCode-In-Go/ListNode"
)

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next //偶数，特殊处理
	}
	cut(head, slow)
	return isEqual(head, reverse(slow))
}

func cut(head *ListNode, cutNode *ListNode) {
	for head.Next != cutNode {
		head = head.Next
	}
	head.Next = nil
}

func reverse(head *ListNode) *ListNode {
	newHead :=&ListNode{}
	for head != nil {
		nextNode := head.Next
		head.Next = newHead
		newHead = head
		head = nextNode
	}
	return newHead
}

func isEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}