/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (31.67%)
 * Likes:    814
 * Dislikes: 1432
 * Total Accepted:    389.1K
 * Total Submissions: 1.2M
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 * 
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 * 
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 * 
 * Example 1:
 * 
 * 
 * Input: 4
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since 
 * the decimal part is truncated, 2 is returned.
 * 
 * 
 */
func mySqrt(x int) int {

	// 앞에서 부터 올려 가면 시간이 너무 걸린다.
	// 바이너리 서치를 해야 겠다!!!
	for i:= 1; i <= x; {
		if i*i > x {
			return i-1
		}
	}

	if x == 1 {
		return 1
	}

	return 0
    
}

