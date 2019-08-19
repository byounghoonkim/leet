/*
 * @lc app=leetcode id=20 lang=c
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (35.98%)
 * Total Accepted:    528.4K
 * Total Submissions: 1.5M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 * 
 * An input string is valid if:
 * 
 * 
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * 
 * 
 * Note that an empty string is also considered valid.
 * 
 * Example 1:
 * 
 * 
 * Input: "()"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "()[]{}"
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "(]"
 * Output: false
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: "([)]"
 * Output: false
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: "{[]}"
 * Output: true
 * 
 * 
 */

struct Node {
	char c;
	struct Node* next;
};

struct Node* head;

void push(char c) {
	struct Node* l = malloc(sizeof(struct Node));
	l->next = head; 
	l->c = c; 
	head = l;
}

char pop() {
	if(NULL == head) {
		return '\0';
	}

	char c = head->c;
	struct Node* tmp = head->next;
	free(head);
	head = tmp;

	return c;
}

bool isValid(char* s) {
	head = NULL;

	int i = 0;
	while(s[i] != '\0') {
		char c = s[i];
		if(c == '(') push(c);
		if(c == '[') push(c);
		if(c == '{') push(c);

		if(c == ')'){
			if('(' != pop()) return false;
		}

		if(c == ']'){
			if('[' != pop()) return false;
		}

		if(c == '}'){
			if('{' != pop()) return false;
		}

		i++;
	}

	if('\0' != pop())
		return false;

	return true;
}
