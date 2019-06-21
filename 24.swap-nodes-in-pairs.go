/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		if cur.Next != nil {
			temp := cur.Val
			cur.Val = cur.Next.Val
			cur.Next.Val = temp
			cur = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}
	return head
}

