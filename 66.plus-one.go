/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 *
 * https://leetcode.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (41.44%)
 * Likes:    920
 * Dislikes: 1628
 * Total Accepted:    413.7K
 * Total Submissions: 998.1K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty array of digits representing a non-negative integer, plus
 * one to the integer.
 * 
 * The digits are stored such that the most significant digit is at the head of
 * the list, and each element in the array contain a single digit.
 * 
 * You may assume the integer does not contain any leading zero, except the
 * number 0 itself.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,3]
 * Output: [1,2,4]
 * Explanation: The array represents the integer 123.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [4,3,2,1]
 * Output: [4,3,2,2]
 * Explanation: The array represents the integer 4321.
 * 
 */
func plusOne(digits []int) []int {
	l := len(digits)
	digits[l-1] += 1
	for i:= l-1; i >0 ; i-- {
		if digits[i] == 10 {
			digits[i] = 0 
			digits[i-1] += 1
		}
	}

	
	if digits[0] == 10 {
		digits[0] = 0 
		digits = append([]int{1}, digits...)
	}

	return digits
    
}

