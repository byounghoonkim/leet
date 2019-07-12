/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.27%)
 * Likes:    392
 * Dislikes: 1647
 * Total Accepted:    278.4K
 * Total Submissions: 862.7K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 * 
 * If the last word does not exist, return 0.
 * 
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 * 
 * Example:
 * 
 * 
 * Input: "Hello World"
 * Output: 5
 * 
 * 
 * 
 * 
 */
func lengthOfLastWord(s string) int {
	ret := 0
	temp := 0
	for _, c:= range(s) {
		if c == ' ' {
			if temp != 0 {
				ret = temp
				temp = 0
			} 
		} else {
			temp ++
		}
	}

	if temp != 0 {
		ret = temp
	}

	return ret
}

