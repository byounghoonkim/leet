/*
 * @lc app=leetcode id=89 lang=golang
 *
 * [89] Gray Code
 */

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	ret := []int{0, 1}
	for i := 1; i < n; i++ {
		l := len(ret)
		for j := l - 1; j >= 0; j-- {
			ret = append(ret, ret[j]+1<<uint(i))
		}
	}

	return ret
}

