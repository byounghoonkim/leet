/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {

	stack := make([]int, 0, k)
	cur := head
	start := cur

	for cur != nil {
		stack = append(stack, cur.Val)

		if len(stack) == k {
			for len(stack) != 0 {
				var x int
				x, stack = stack[len(stack)-1], stack[:len(stack)-1]
				start.Val = x
				start = start.Next
			}
		}
		cur = cur.Next
	}

	return head
}

