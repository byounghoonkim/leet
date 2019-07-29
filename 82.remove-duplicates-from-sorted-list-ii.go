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


//  next 메소드는 파라미터로 넘어온 노드를 기준으로 다음 항목을 리턴해 준다.
// 파라미터의 노드 가 닐이거나 다음 노드가 닐이 일때는 자기 자신을 리턴한다.
// 노드가 다음 노드와 값이 다를 때도 자기 자신을 리턴한다.
// 그렇지 않은 경우는 현재 노드와 다음 노드의 값이 일치 하는 경우이므로
// 다음 노드가 닐이거나 일치 하지 않을 때까지 진행한다.
// 그리고 얻어진 노드를 (이 노드도 중복 값이 배열되어 있을 수 있으므로)
// 재귀 호출하여 결과를 리턴한다.
func next(cur *ListNode) *ListNode {
	if cur == nil { return nil }
	if cur.Next == nil { return cur}
	if cur.Next.Val != cur.Val { return cur }
	for  cur.Next != nil {
		if cur.Next.Val != cur.Val {
			break
		}
		cur = cur.Next
	}

	return next(cur.Next)
}
func deleteDuplicates(head *ListNode) *ListNode {
	ret := next(head)
	if ret == nil { return ret }

	cur := ret
	for ;;{
		n := next(cur.Next)
		cur.Next = n
		if n == nil {
			break
		}
		cur = n
	}

	return ret
}

