/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 *
 * https://leetcode.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (49.04%)
 * Likes:    2089
 * Dislikes: 65
 * Total Accepted:    354.3K
 * Total Submissions: 722.4K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given a set of candidate numbers (candidates) (without duplicates) and a
 * target number (target), find all unique combinations in candidates where the
 * candidate numbers sums to target.
 *
 * The same repeated number may be chosen from candidates unlimited number of
 * times.
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
 * Input: candidates = [2,3,6,7], target = 7,
 * A solution set is:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,3,5], target = 8,
 * A solution set is:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 */

func combi(candidates []int, target int, path []int, ret *[][]int) {
	// 전달된 후보 슬라이스 길이가 0 이면 리턴한다
	if len(candidates) == 0 {
		return
	}

	// target 값이 0 이하이면 리턴한다
	if target < 0 {
		return
	}

	// 전달된 후보 값 첫번째를 꺼내서
	n := candidates[0]
	// 첫번째 값의 배수 스텝으로 루프를 수행하는데 target 값 이하일때까지만 수행한다
	for i := 0; i <= target; i += n {
		// i가 0이면 첫번째 요소를 쓰지 않는 패스 이므로
		// path 에 추가하지 않는다.
		if i != 0 {
			path = append(path, n)
		}

		// 첫후보값의 배수가 타겟 값과 일치 했을 때 현재 경로를
		// 복사 해 결과 슬라이스에 넣어 준다.
		if target == i {
			t := make([]int, len(path))
			copy(t, path)
			*ret = append(*ret, t)
		} else {
			// 일치하지 않을 때는 첫번째 후보값을 제외한 후보값 슬라이스와
			// 타겟에서 현재 i를 뺀 값, 그리고 현재 까지 경로를
			// 가지고 재귀 호출한다.
			combi(candidates[1:], target-i, path, ret)
		}
	}

}

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	path := []int{}
	combi(candidates, target, path, &ret)
	return ret

	// 이 알고리즘은 후보값 개수 n 에 루프를 순회 하면서 각 요소가
	// 타겟 값 내에서 배수 만큼 루프를 추가로 돈다.
	// n * (t / ni) 정도의 실행 횟수를 가진다.
}

