package main

import (
	. "LeetCode-In-Go/ListNode"
)

/**
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := getLength(headA)
	lenB := getLength(headB)
	for{
		if lenA > lenB {
			headA = headA.Next
			lenA--
		}else if lenA < lenB {
			headB = headB.Next
			lenB--
		}else {
				break
		}
	}
	var node *ListNode
	for i:= 0; i < lenA; i++ {
		if headA.Val == headB.Val {
			node = headA
			break;
		}else{
			headA = headA.Next
			headB = headB.Next
		}
	}
	return node
}

func getLength(headA *ListNode) int {
	len := 0
	for headA.Next != nil {
		len++
		headA = headA.Next
	}
	return len
}

 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	lenA,lenB := 0, 0
	for a != nil {
		lenA++
		a = a.Next
	}
	for b != nil {
		lenB++
		b = b.Next
	}
	a, b = headA, headB
	for lenA > lenB {
		a = a.Next
		lenA--
	}
	for lenB > lenA {
		b = b.Next
		lenB--
	}
	for a != b {
		a = a.Next
		b = b.Next
	}
	return a
}