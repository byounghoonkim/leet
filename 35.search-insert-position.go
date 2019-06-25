/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 *
 * https://leetcode.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (40.93%)
 * Likes:    1349
 * Dislikes: 181
 * Total Accepted:    412.5K
 * Total Submissions: 1M
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * Given a sorted array and a target value, return the index if the target is
 * found. If not, return the index where it would be if it were inserted in
 * order.
 *
 * You may assume no duplicates in the array.
 *
 * Example 1:
 *
 *
 * Input: [1,3,5,6], 5
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [1,3,5,6], 2
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: [1,3,5,6], 7
 * Output: 4
 *
 *
 * Example 4:
 *
 *
 * Input: [1,3,5,6], 0
 * Output: 0
 *
 *
 */

func searchBinary(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		if nums[l] == target {
			return l
		}

		if nums[r] == target {
			return r
		}

		// 값을 찾았으면 리턴하지만
		// 못 찾을 때는 두 인덱스가 인접할 때까지 바이너리 서치를 한다.
		// 두 인덱스가 인접하면 오른쪽 인덱스를 리턴한다.
		if l+1 == r {
			return r
		}

		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			l = m
		} else {
			r = m
		}
	}

	return -1
}
func searchInsert(nums []int, target int) int {

	// 첫번째 값 보다 타겟이 작으면 0을 리턴한다.
	if nums[0] > target {
		return 0
	}

	// 마지막 값 보다 타겟이 크면 마지막 인덱스 다음 값을 리턴한다.
	if nums[len(nums)-1] < target {
		return len(nums)
	}

	//여기까지 왔으면 배열의 범위 안에 타겟이 있다는 것이 보장 된다.
	return searchBinary(nums, target)

	// 이 알고리즘은 주어진 nums 배열의 크기 n 에 대해 바이너리 서치를 진행 하므로
	// 바이너리 트리 깊이 만큼인 logn 만큼 실행 시간 복잡도를 가진다.
	// O(logn)
}

