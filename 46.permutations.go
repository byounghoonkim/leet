/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

func per(pr *[][]int, nums []int, depth int) {
	if depth == len(nums) {
		a := make([]int, len(nums))
		copy(a, nums)
		*pr = append(*pr, a)
	}

	for i := depth; i < len(nums); i++ {
		nums[i], nums[depth] = nums[depth], nums[i]
		per(pr, nums, depth+1)
		nums[i], nums[depth] = nums[depth], nums[i]
	}
}

func permute(nums []int) [][]int {
	var ret [][]int

	per(&ret, nums, 0)

	return ret
}
