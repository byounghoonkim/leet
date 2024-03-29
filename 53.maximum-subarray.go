
/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (43.82%)
 * Likes:    4472
 * Dislikes: 156
 * Total Accepted:    555.2K
 * Total Submissions: 1.3M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */

func maxSubArray(nums []int) int {

	ret := nums[0]

	for i, _ := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if ret < sum {
				ret = sum
			}
		}
	}

	return ret

}
