/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */
func removeElement(nums []int, val int) int {
	var found int
	l := len(nums)

	for i, _ := range nums {

		if i >= l - found {
			break
		} 

		for (nums[i] == val) && (i < l-found) {
			found += 1
			nums[i], nums[l-found] = nums[l-found], nums[i]  // swap
		}

	}
	
	return l - found
}

