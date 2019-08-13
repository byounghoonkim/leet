/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 *
 * https://leetcode.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (35.54%)
 * Likes:    1364
 * Dislikes: 97
 * Total Accepted:    207.1K
 * Total Submissions: 580.8K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * Reverse a linked list from position m to n. Do it in one-pass.
 *
 * Note: 1 ≤ m ≤ n ≤ length of list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL, m = 2, n = 4
 * Output: 1->4->3->2->5->NULL
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(head *ListNode, c int) *ListNode {

	remain := head
	var newHead *ListNode
	var newTail *ListNode

	for i := 0; i < c; i++ {
		temp := remain
		remain = remain.Next

		if newHead == nil {
			newHead = temp
			newTail = temp
			newTail.Next = nil
		} else {
			temp.Next = newHead
			newHead = temp
		}
	}

	if newTail != nil {
		newTail.Next = remain
	}

	if newHead != nil {
		return newHead
	}

	return head

}

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	var headLast *ListNode
	var revHead *ListNode
	for i := 1; i < m; i++ {
		if i == 1 {
			headLast = head
		} else {
			headLast = headLast.Next
		}
	}

	if headLast == nil {
		revHead = head
	} else {
		revHead = headLast.Next
	}

	revHead = reverse(revHead, n-m+1)

	if headLast == nil {
		return revHead
	} else {
		headLast.Next = revHead
		return head
	}
}

