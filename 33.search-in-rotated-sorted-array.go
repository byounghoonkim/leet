/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (32.96%)
 * Likes:    2506
 * Dislikes: 317
 * Total Accepted:    428.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 *
 * You are given a target value to search. If found in the array return its
 * index, otherwise return -1.
 *
 * You may assume no duplicate exists in the array.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * Example 1:
 *
 *
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 *
 */
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		if nums[l] == target {
			return l
		}
		if nums[r] == target {
			return r
		}

		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

		// 이 알고리즘의 핵심이다. nums[m] 값이 nums [l] 보다 크다면
		// 		m 의 왼쪽이 정렬 되어 있으므로 해당 구간에 target 이 속하는지 확인한다.
		// 		속하지 않으면 정렬되지 않은 오른쪽으로 탐색을 계속한다.
		// nums[m] 값이 nums [l] 보다 작다면
		// 		m 의 오른쪽이 이 정렬되어 있으므로 이 구간에 targe이 있는지 확인한다.
		// 		없다면 왼쪽 구간에 탐색을 계속한다.
		if nums[l] < nums[m] {
			if nums[l] < target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[m] < target && target < nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return -1

	// 이 알고리즘은 nums 의 개수 n 에 대해 이진 탐색을 진행하므로
	// O(logn) 의 수행 시간 복잡도를 가진다.
}

