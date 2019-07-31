/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 이 솔루션은 82번 문제를 먼저 풀어서 해당 솔루션에
//  next 메소드의 리턴값만 조금 변형하여 제출한다.

func next(cur *ListNode) *ListNode {
	if cur == nil {
		return nil
	}
	if cur.Next == nil {
		return cur
	}
	if cur.Next.Val != cur.Val {
		return cur
	}
	for cur.Next != nil {
		if cur.Next.Val != cur.Val {
			break
		}
		cur = cur.Next
	}

	return cur // solution 82 : return next(cur.Next)
}
func deleteDuplicates(head *ListNode) *ListNode {
	ret := next(head)
	if ret == nil {
		return ret
	}

	cur := ret
	for {
		n := next(cur.Next)
		cur.Next = n
		if n == nil {
			break
		}
		cur = n
	}

	return ret
}

