/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (27.74%)
 * Likes:    4112
 * Dislikes: 379
 * Total Accepted:    624.9K
 * Total Submissions: 2.3M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: "cbbd"
 * Output: "bb"
 *
 *
 */

func isPanlindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
func longestPalindrome(s string) string {
	max := 0
	ret := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			sub := s[i : j+1]
			if isPanlindrome(sub) {
				if max < len(sub) {
					max = len(sub)
					ret = sub
				}
			}
		}
	}

	return ret

}

