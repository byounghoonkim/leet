/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (39.74%)
 * Likes:    1039
 * Dislikes: 198
 * Total Accepted:    317.5K
 * Total Submissions: 798.9K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 * 
 * The input strings are both non-empty and contains only characters 1 or 0.
 * 
 * Example 1:
 * 
 * 
 * Input: a = "11", b = "1"
 * Output: "100"
 * 
 * Example 2:
 * 
 * 
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 * 
 */
func addBinary(a string, b string) string {
	al:= len(a)
	bl:= len(b)
	ai := al -1
	bi := bl -1

	ret := ""
	up := false // 올림을 나타냄
	for ; ai >=0 || bi >= 0 ; {
		n := 0
		if ai >= 0 {
			if a[ai] == '1' {
				n++
			}
		}
		if bi >= 0 {
			if b[bi] == '1' {
				n++
			}
		}
		if up {
			n++
		}

		// 1이거나 3이면 1로 표기
		if 1 == n%2 {
			ret = "1" + ret
		} else {
			ret = "0" + ret
		}

		// 2 , 3 이면 올림 자리를 true 로 만듬
		if 1 == n / 2 {
			up = true
		} else {
			up = false
		}

		ai--
		bi--
	}

	// 마지막 자리에 남아 있는 올림이 있으면 처리 
	if up {
		ret = "1" + ret
	}

	return ret
}

