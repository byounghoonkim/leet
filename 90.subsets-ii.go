import "sort"

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

type Result [][]int

func (r *Result) helper(nums []int, temp []int) {

	if 0 == len(nums) {
		return
	}

	for i, n := range nums {

		if i > 0 && nums[i-1] == n {
			continue
		}

		rtemp := append(temp, n)
		rclone := make([]int, len(rtemp))
		copy(rclone, rtemp)
		*r = append(*r, rclone)
		r.helper(nums[i+1:], rtemp)
	}

}

func subsetsWithDup(nums []int) [][]int {

	ret := make(Result, 1)

	sort.Ints(nums)

	ret.helper(nums, []int{})

	return ret
}

