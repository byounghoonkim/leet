/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 *
 * https://leetcode.com/problems/subsets/description/
 *
 * algorithms
 * Medium (53.89%)
 * Likes:    2122
 * Dislikes: 53
 * Total Accepted:    388.4K
 * Total Submissions: 720.7K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a set of distinct integers, nums, return all possible subsets (the
 * power set).
 * 
 * Note: The solution set must not contain duplicate subsets.
 * 
 * Example:
 * 
 * 
 * Input: nums = [1,2,3]
 * Output:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 * 
 */

type Result [][]int

func (r *Result)helper(nums []int, temp []int)  {

	if 0 == len(nums) {
		return 
	}

	for i, n := range(nums) {
		r := append(temp, n)
		rclone := make([]int, len(r))
		copy(rclone, r)
		*r = append(*r, rclone)
		r.helper(nums[i+1:], r)
	}
	
}

func subsets(nums []int) [][]int {

	ret := make(Result, 1) 

	ret.helper(nums, []int{})

	return ret
    
}

