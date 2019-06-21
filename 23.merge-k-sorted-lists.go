/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {

	var ret *ListNode
	var last *ListNode
	var candidate *ListNode
	var candidate_index int

	is_done := false

	for is_done == false {
		for i, n := range lists {
			if n == nil {
				continue
			}

			if candidate == nil {
				candidate = n
				candidate_index = i
			} else {
				if candidate.Val > n.Val {
					candidate = n
					candidate_index = i
				}
			}
		}

		if candidate != nil {
			if last != nil {
				last.Next = lists[candidate_index]
				lists[candidate_index] = lists[candidate_index].Next
				last = last.Next
				last.Next = nil
			} else {
				last = lists[candidate_index]
				lists[candidate_index] = lists[candidate_index].Next
				last.Next = nil
				ret = last
			}
			candidate = nil
		} else {
			is_done = true
		}
	}

	return ret

}

