package main

import (
	. "LeetCode-In-Go/ListNode"
)
//author: hawchou
func removeNthFromEnd(head *ListNode, n int) *ListNode{
	fast := head
	for n > 0 {
		fast = fast.Next
		n--
	}
	if fast == nil {
		return head.Next
	}
	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head;
}
