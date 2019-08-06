/*
 * @lc app=leetcode id=87 lang=golang
 *
 * [87] Scramble String
 */

func checkScramble(s1, s2 string, cache map[string]bool) bool {
	n := len(s1)
	if n == 1 {
		return s1 == s2
	}

	if t, ok := cache[s1+" "+s2]; ok {
		return t
	}

	for i := 1; i < n; i++ {
		match := (checkScramble(s1[:i], s2[:i], cache) && checkScramble(s1[i:], s2[i:], cache)) ||
			(checkScramble(s1[:i], s2[n-i:], cache) && checkScramble(s1[i:], s2[:n-i], cache))

		if match {
			cache[s1+" "+s2] = true
			return true
		}
	}

	cache[s1+" "+s2] = false

	return false
}

func isScramble(s1 string, s2 string) bool {
	cache := make(map[string]bool)
	return checkScramble(s1, s2, cache)
}

