/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (48.62%)
 * Likes:    833
 * Dislikes: 48
 * Total Accepted:    213.3K
 * Total Submissions: 438.6K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 * 
 * Example:
 * 
 * 
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 * 
 * 
 */

func helper(n int, s int , k int, temp []int, ret *[][]int) {
	if k == 0 {
		t := make([]int, len(temp))
		copy(t, temp)
		*ret = append(*ret, t)
		return
	}

	for i := s ; i <= n  ; i++{
		helper(n, i+1, k-1, append(temp, i), ret)
	}
}

func combine(n int, k int) [][]int {
	ret := make([][]int, 0) 
	helper(n, 1, k, make([]int, 0), &ret)
	return ret
}

