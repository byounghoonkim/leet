/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

func per(pr *[][]int, nums []int, depth int) {
	if depth == len(nums) {
		a := make([]int, len(nums))
		copy(a, nums)

		*pr = append(*pr, a)
		return
	}

	visited := make(map[int]bool)
	for i := depth; i < len(nums); i++ {
		if v := visited[nums[i]]; v == true {
			continue
		}

		nums[i], nums[depth] = nums[depth], nums[i]
		per(pr, nums, depth+1)
		nums[i], nums[depth] = nums[depth], nums[i]

		visited[nums[i]] = true

	}
}

func permuteUnique(nums []int) [][]int {
	var ret [][]int

	per(&ret, nums, 0)

	return ret
}
