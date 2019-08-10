/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */
func numDecodings(s string) int {
	return countWay(s)
}

func canDecode(s string) bool {
	if len(s) == 1 {
		if s[0] >= '1' && s[0] <= '9' {
			// 1 ... 9
			return true
		}
	}

	if len(s) == 2 {
		if s[0] == '1' {
			if s[1] >= '0' && s[1] <= '9' {
				// 10 ... 19
				return true
			}
		}

		if s[0] == '2' {
			if s[1] >= '0' && s[1] <= '6' {
				// 20 ... 26
				return true
			}
		}
	}

	return false
}

func countWay(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		if canDecode(s) {
			return 1
		} else {
			return 0
		}
	}

	ret := 0

	if canDecode(s[:1]) {
		ret += countWay(s[1:])
	}

	if len(s) >= 2 && canDecode(s[:2]) {
		if len(s) == 2 {
			ret += 1
		} else {
			ret += countWay(s[2:])
		}
	}

	return ret
}

