/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */
func strStr(haystack string, needle string) int {
	nl := len(needle)
	hl := len(haystack)

	if nl == 0 {
		return 0
	}

	for i := 0 ; i < hl ; i++ {
		if nl > hl - i {
			break
		}

		found := true
		for j := 0 ; j < nl ; j++ {
			if needle[j] != haystack[i+j] {
				found = false
				break
			}
		}

		if found {
			return i
		}
	}

	return -1
}

