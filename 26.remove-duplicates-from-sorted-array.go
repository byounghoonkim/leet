/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
func removeDuplicates(nums []int) int {
	var last int 
	var dup_count int

	for i, n := range nums {
		fmt.Println(i, n)
		if i == 0 {
			last = 0
			continue
		}
		if nums[last] == n {
			dup_count += 1
		} else {
			nums[last+1] = n
			last += 1 
		}
	}

	return len(nums) - dup_count
    
}

