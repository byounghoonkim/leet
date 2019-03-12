/*
 * @lc app=leetcode id=21 lang=c
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (46.03%)
 * Total Accepted:    522.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 * 
 * Example:
 * 
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 * 
 * 
 */
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2) {

	struct ListNode* head = NULL;
	struct ListNode* last = NULL;
	struct ListNode* target = NULL;
	struct ListNode* t1;
	struct ListNode* t2;

	t1 = l1;
	t2 = l2;

	while( t1 != NULL || t2 != NULL ) {
		target = NULL;

		if(t1 == NULL) {
			target = t2;
			t2 = t2->next;
		} else {
			if(t2 == NULL) {
				target = t1;
				t1 = t1->next;
			}
		}

		if(NULL == target)
		{
			if(t1->val > t2->val)
			{
				target = t2;
				t2 = t2->next;
			} else {
				target = t1;
				t1 = t1->next;
			}
		}

		if(NULL != last) {
			last->next = target;
		}

		last = target;
		last->next = NULL;

		if(NULL == head){
			head = last;
		} 
			
	}

	return head;
    
}
