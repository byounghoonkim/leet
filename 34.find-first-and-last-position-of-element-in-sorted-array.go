/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (33.71%)
 * Likes:    1638
 * Dislikes: 87
 * Total Accepted:    310.7K
 * Total Submissions: 921.9K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * If the target is not found in the array, return [-1, -1].
 *
 * Example 1:
 *
 *
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 *
 * Example 2:
 *
 *
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 *
 */

func searchBinaryVjjjj(nums []int, target int) int {
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

		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	// 이진 탐색을 해 target 을 찾아낸다.
	i := searchBinary(nums, target)

	// 찾아 내지 못했다면 -1로 리턴한다.
	if i == -1 {
		return []int{i, i}
	}

	s := i
	e := i

	// i 앞의 값이 target 가 일치 하면 s 를 구하기 위해 앞으로 전진하며
	// target과 다른 값의 다음 인덱스를 s에 저장 되도록 한다.
	for s = i; s > 0; s-- {
		if nums[s-1] != target {
			break
		}
	}

	for e = i; e < len(nums)-1; e++ {
		if nums[e+1] != target {
			break
		}
	}

	return []int{s, e}

	// 이 알고리즘은 이진 탐색을 위해 O(logn) 의 수행 시간에 더해
	// 같은 값의 처음 끝을 찾아 내기 위해 최악의 경우 nums의 길이 만큼 탐색을 해야 하므로
	// O(logn) + O(n) 정도의 실행 시간 복잡도를 가진다.
}

