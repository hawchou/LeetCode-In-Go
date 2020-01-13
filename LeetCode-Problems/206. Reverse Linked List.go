package main

import (
	. "LeetCode-In-Go/ListNode"
)
/**
approach 1: 头插法
 */
func reverseList(head *ListNode) *ListNode {
	r := &ListNode{Val:-1}
	for head != nil {
		next := head.Next
		head.Next = r.Next
		r.Next = head
		head = next
	}
	return r.Next
}