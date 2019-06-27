/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 *
 * https://leetcode.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (29.07%)
 * Likes:    1744
 * Dislikes: 587
 * Total Accepted:    219.3K
 * Total Submissions: 754.4K
 * Testcase Example:  '[1,2,0]'
 *
 * Given an unsorted integer array, find the smallest missing positive
 * integer.
 *
 * Example 1:
 *
 *
 * Input: [1,2,0]
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: [3,4,-1,1]
 * Output: 2
 *
 *
 * Example 3:
 *
 *
 * Input: [7,8,9,11,12]
 * Output: 1
 *
 *
 * Note:
 *
 * Your algorithm should run in O(n) time and uses constant extra space.
 *
 */

func firstMissingPositive(nums []int) int {
	ret := 0
	// int - bool 맵을 하나 만든다.
	a := make(map[int]bool)

	// 각 값을 순회 하면서 0 보다 큰 값을 맵에 추가한다.
	for _, n := range nums {
		if n > 0 {
			a[n] = true
		}
	}

	// 1부터 무한대로 커지면서
	// 맵을 찾되 처음 찾지 못하면 바로 리턴한다. 그 값이 정답이다.
	for i := 1; ; i++ {
		if false == a[i] {
			ret = i
			break
		}
	}

	return ret
}

