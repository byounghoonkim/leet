/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 *
 * https://leetcode.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (42.04%)
 * Likes:    864
 * Dislikes: 47
 * Total Accepted:    227.2K
 * Total Submissions: 540.3K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * Given a collection of candidate numbers (candidates) and a target number
 * (target), find all unique combinations in candidates where the candidate
 * numbers sums to target.
 *
 * Each number in candidates may only be used once in the combination.
 *
 * Note:
 *
 *
 * All numbers (including target) will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [10,1,2,7,6,1,5], target = 8,
 * A solution set is:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,5,2,1,2], target = 5,
 * A solution set is:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 */

func combi(candidates []int, target int, path []int, ret *[][]int) {
	if len(candidates) == 0 {
		return
	}

	if target < 0 {
		return
	}

	n := candidates[0]

	// 39번 문제와 유사 하지만 다른 점은
	// 동일한 후보값을 갯수를 세어서
	// 해당 값의 갯수에 맞춘 조합을 만들어 낸다는 것이 다르다
	// 만약 2가 3개 있다면
	// [] [2] [22] [222] 와 같은 조합을 만들어 내려고 한다.
	count := 0
	for _, a := range candidates {
		if a == n {
			count += 1
		} else {
			break
		}
	}

	// 위에서 구한 갯수 만큼 루프를 돈다.
	for i := 0; i <= count; i++ {
		if i != 0 {
			path = append(path, n)
		}

		if target == i*n {
			t := make([]int, len(path))
			copy(t, path)
			*ret = append(*ret, t)
		} else {
			combi(candidates[count:], target-i*n, path, ret)
		}
	}

}

func sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	ret := [][]int{}
	path := []int{}

	// 각 후보 값을 같은 값끼리 묶기 위해 먼저 정렬을 한다.
	sort(candidates)

	combi(candidates, target, path, &ret)
	return ret
}

