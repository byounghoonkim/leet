/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func next(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	if head.Val != head.Next.Val {
		return head
	}

	cur := head
	for  cur.Next != nil || cur.Next.Val == cur.Val {
		cur = cur.Next
	}

	if cur.Next == nil {
		return nil
	}

	return next(cur.Next.Next)
}
func deleteDuplicates(head *ListNode) *ListNode {

	ret := next(head)

	cur := ret
	for n := next(cur.Next) ; n != nil; {
		fmt.Println("cur", cur)
		fmt.Println("cur.Next", cur.Next)
		fmt.Println("n", n)
		fmt.Println("||||")
		cur.Next = n
		cur = cur.Next
	}

	return ret
}

