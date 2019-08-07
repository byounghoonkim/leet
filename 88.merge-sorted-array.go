/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
func merge(nums1 []int, m int, nums2 []int, n int) {

	for i := m + n - 1; i >= 0; i-- {

		if m-1 >= 0 && n-1 >= 0 {
			if nums1[m-1] > nums2[n-1] {
				nums1[i] = nums1[m-1]
				m--
			} else {
				nums1[i] = nums2[n-1]
				n--
			}
		} else {
			if m-1 < 0 {
				nums1[i] = nums2[n-1]
				n--
				continue
			}
			if n-1 < 0 {
				nums1[i] = nums1[m-1]
				m--
				continue
			}
		}
	}
}

