import "fmt"

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (40.52%)
 * Total Accepted:    352.2K
 * Total Submissions: 869.1K
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent.
 *
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 *
 * Example:
 *
 *
 * Input: "23"
 * Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * Note:
 *
 * Although the above answer is in lexicographical order, your answer could be
 * in any order you want.
 *
 */

func letterCombinations(digits string) []string {

	if 0 == len(digits) {
		return make([]string, 0)
	}

	t := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	count := 1
	for _, n := range digits {
		count *= len(t[string(n)])
	}

	ret := make([]string, count)

	step := count
	fmt.Print(step)
	for _, n := range digits {
		step = step / len(t[string(n)])
		for j, r := range ret {
			ret[j] = r + string(t[string(n)][j/step%len(t[string(n)])])
		}
	}

	return ret
}
