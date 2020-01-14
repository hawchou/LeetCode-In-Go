package main

import (
	. "LeetCode-In-Go/ListNode"
)

func swapPairs(head *ListNode) *ListNode {
	node := &ListNode{Val:-1}
	node.Next = head;
	pre := node
	for pre.Next != nil && pre.Next.Next != nil {
		l1, l2 := pre.Next, pre.Next.Next
		next := l2.Next
		l1.Next = next
		l2.Next = l1
		pre.Next = l2

		pre = l1
	}
	return node.Next
}