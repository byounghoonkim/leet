/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 *
 * https://leetcode.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (32.08%)
 * Likes:    2048
 * Dislikes: 204
 * Total Accepted:    277K
 * Total Submissions: 861.7K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * Given an array of non-negative integers, you are initially positioned at the
 * first index of the array.
 * 
 * Each element in the array represents your maximum jump length at that
 * position.
 * 
 * Determine if you are able to reach the last index.
 * 
 * Example 1:
 * 
 * 
 * Input: [2,3,1,1,4]
 * Output: true
 * Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last
 * index.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,2,1,0,4]
 * Output: false
 * Explanation: You will always arrive at index 3 no matter what. Its
 * maximum
 * jump length is 0, which makes it impossible to reach the last index.
 * 
 * 
 */

func jump(nums []int, cur int, cache []bool) bool  {
	if len(nums) -1 == cur {
		return true
	}

	if nums[cur] == 0 {
		return false
	}

	if cache[cur] == true {
		return false
	}

	j := nums[cur]
	if nums[cur] + cur >= len(nums) {
		j = len(nums) -1 -cur
		
	}

	//for k := 1; k <= j ; k++ {
	for k := j; k >= 1 ; k-- {

		if true == jump(nums, cur+k, cache) {
			return true
		} else {
			cache[cur+k] = true 
		}
	}

	cache[cur] = true 
	return false
	
}

func canJump(nums []int) bool {
	cache := make([]bool, len(nums))
	return jump(nums, 0, cache)
    
}

