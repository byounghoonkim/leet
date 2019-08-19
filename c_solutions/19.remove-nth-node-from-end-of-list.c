/*
 * @lc app=leetcode id=19 lang=c
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (34.03%)
 * Total Accepted:    359.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, remove the n-th node from the end of list and return
 * its head.
 * 
 * Example:
 * 
 * 
 * Given linked list: 1->2->3->4->5, and n = 2.
 * 
 * After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 * 
 * 
 * Note:
 * 
 * Given n will always be valid.
 * 
 * Follow up:
 * 
 * Could you do this in one pass?
 * 
 */
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* removeNthFromEnd(struct ListNode* head, int n) {
	int len = 0;
	struct ListNode* cur;

	struct ListNode* target;
	struct ListNode* pre;
	cur = head;
	while(cur != NULL) {
		len++;
		cur = cur->next;
	}

	target = head;
	len -= n;
	while(len > 0) {
		len--;
		pre = target;
		target = target->next;
	}

	if(target) {
		if(target == head) {
			return head->next;
		} else {
			pre->next = target->next;
		}
	}

	return head;
    
}
