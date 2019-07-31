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

	var cur *ListNode
	var pre *ListNode
	var first *ListNode

	cur = head
	for ; cur != nil ; {
		fmt.Println(cur)
		var next *ListNode
		if cur.Val < x {
			if pre != nil {
				pre.Next = cur.Next
			} 

			next = cur.Next
			if first != nil {
				cur.Next = first.Next
				first.Next = cur
			}  else {
				first = cur
			}
			cur = next
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	fmt.Println(first)

	return first
}

