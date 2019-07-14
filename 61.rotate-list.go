/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 *
 * https://leetcode.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (27.44%)
 * Likes:    644
 * Dislikes: 828
 * Total Accepted:    200K
 * Total Submissions: 728.6K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, rotate the list to the right by k places, where k is
 * non-negative.
 * 
 * Example 1:
 * 
 * 
 * Input: 1->2->3->4->5->NULL, k = 2
 * Output: 4->5->1->2->3->NULL
 * Explanation:
 * rotate 1 steps to the right: 5->1->2->3->4->NULL
 * rotate 2 steps to the right: 4->5->1->2->3->NULL
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 0->1->2->NULL, k = 4
 * Output: 2->0->1->NULL
 * Explanation:
 * rotate 1 steps to the right: 2->0->1->NULL
 * rotate 2 steps to the right: 1->2->0->NULL
 * rotate 3 steps to the right: 0->1->2->NULL
 * rotate 4 steps to the right: 2->0->1->NULL
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {

	// 리스트의 길이를 구한다.
	l := 0 
	cur := head
	for cur != nil {
		l++
		cur = cur.Next
	}

	// 길이가 0이나 1이면 아무리 돌려도 제자리다. 그냥 리턴한다.
	if l <= 1 {
		return head
	}

	// 로테이트 한 횟수가 길이의 정배수 이면 
	// 제자리이므로 원본 리스트를 리턴한다.
	if 0 == k % l {
		return head
	}


	// l - k 번째 요소를 찾는다.
	cur = head
	count := 0
	var newTail *ListNode
	var newHead *ListNode
	for cur != nil {
		count++
		if count >= l - (k%l) {
			//찾은 요소를 newTail로 저장하고
			// 다음 항목을 newHead로 저장한다.
			newTail = cur
			newHead = newTail.Next
			newTail.Next =nil 
			break
		}
		cur = cur.Next
	}

	// newHead 에 마지막 요소를 찾아 head 로 연결한다.
	cur = newHead 
	for cur != nil {
		if cur.Next == nil {
			cur.Next = head
			break
		}

		cur = cur.Next
	}

	return newHead

}

