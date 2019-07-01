/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (44.28%)
 * Likes:    11072
 * Dislikes: 376
 * Total Accepted:    1.9M
 * Total Submissions: 4.3M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 * 
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 * 
 * Example:
 * 
 * 
 * Given nums = [2, 7, 11, 15], target = 9,
 * 
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 * 
 * 
 */
func twoSum(nums []int, target int) []int {

	for i, _:= range nums {
		for j:= i+1; j < len(nums);j++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return []int{}
}

