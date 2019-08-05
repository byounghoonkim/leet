/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	var first *ListNode
	var first_last *ListNode
	var second *ListNode
	var second_last *ListNode

	cur := head
	for cur != nil {

		if cur.Val < x {
			if first == nil {
				first = cur
				first_last = cur
			} else {
				first_last.Next = cur
				first_last = cur
			}
		} else {
			if second == nil {
				second = cur
				second_last = cur
			} else {
				second_last.Next = cur
				second_last = cur
			}
		}

		temp := cur.Next
		cur.Next = nil
		cur = temp
	}

	if first_last != nil {
		first_last.Next = second
	} else {
		if second != nil {
			first = second
		}
	}

	return first
}

