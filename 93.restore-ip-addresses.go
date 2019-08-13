import "strconv"

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 *
 * https://leetcode.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (32.00%)
 * Likes:    728
 * Dislikes: 304
 * Total Accepted:    148K
 * Total Submissions: 460.8K
 * Testcase Example:  '"25525511135"'
 *
 * Given a string containing only digits, restore it by returning all possible
 * valid IP address combinations.
 *
 * Example:
 *
 *
 * Input: "25525511135"
 * Output: ["255.255.11.135", "255.255.111.35"]
 *
 *
 */

func isValidNum(s string) bool {
	i, _ := strconv.Atoi(s)

	// 2자 이상인데 0으로 시작하면 잘못된 숫자로 판단한다.
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	// 0 부터 255 까지를 유효한 숫자로 판단한다.
	if i >= 0 && i <= 255 {
		return true
	}
	return false
}

func restore(s, temp string, depth int, ips *[]string) {
	if depth == 4 {
		if len(s) > 0 && isValidNum(s) {
			*ips = append(*ips, temp+"."+s)
		}
		return
	}

	// 처음을 제외 하고 .을 추가해 준다.
	if depth != 1 {
		temp += "."
	}

	for i := 1; i <= 3; i++ {
		if len(s) >= i && isValidNum(s[:i]) {
			restore(s[i:], temp+s[:i], depth+1, ips)
		}
	}
}

func restoreIpAddresses(s string) []string {
	ret := []string{}
	restore(s, "", 1, &ret)
	return ret
}

