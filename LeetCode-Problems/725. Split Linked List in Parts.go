package main


import (
	. "LeetCode-In-Go/ListNode"
)

// have bug
func splitListToParts(root *ListNode, k int) []*ListNode {
	len := 0
	head := root
	for head.Next != nil {
		head = head.Next;
		len++
	}
	ans := []*ListNode{}
	l, r:= len / k, len % k
	head = root
	prev := &ListNode{}
	for i := 0; i < k; i++ {
		r--
		ans[i] = head
		if r > 0 {
			l = l + 1
		}
		for j := 0; j < l; j++ {
			prev = head
			head = head.Next
		}
		if prev != nil {
			prev.Next = nil
		}
	}
	return ans
}
