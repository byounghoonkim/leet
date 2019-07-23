/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 *
 * https://leetcode.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (31.33%)
 * Likes:    2467
 * Dislikes: 165
 * Total Accepted:    252.5K
 * Total Submissions: 806K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * Given a string S and a string T, find the minimum window in S which will
 * contain all the characters in T in complexity O(n).
 * 
 * Example:
 * 
 * 
 * Input: S = "ADOBECODEBANC", T = "ABC"
 * Output: "BANC"
 * 
 * 
 * Note:
 * 
 * 
 * If there is no such window in S that covers all characters in T, return the
 * empty string "".
 * If there is such window, you are guaranteed that there will always be only
 * one unique minimum window in S.
 * 
 * 

 "ADOBECODEBANC", T = "ABC"
  l
  r

  l--->r
   l   r
   l    --->r
   --->l    r
	   l     ->r
	   -------lr

 */

 func isCover(m map[byte]int) bool {
	 // 커버가 된 것을 판단 하는 것은 맵의 모든 값이 0 이하여야 한다. (0 이하라는 것이 중요하다. 0이 아니라.)
	 // 마이너스 값이 될 수 있는데 t 에 속한 문자가 s 에 다수개 포함되면 마이너스 이다.
	 for _, a := range(m) {
		 if a > 0 { return false }
	 }
	 return true
 }

func minWindow(s string, t string) string {
	l, r := 0,0
	left, right := 0,0

	// target 문자열의 각 문자를 키로 하고 갯수를 값으로 하는 맵을 생성한다.
	m := map[byte]int{}
	for i:=0; i < len(t); i++ {
		m[t[i]] ++
	}

	// 짧은 길이를 기억하는 변수
	min_len := -1

	// 왼쪽 인덱스가 끝까지 진행할 때까지 진행
	for l < len(s)  {
		// 커버가 되거나 오른쪽 인덱스가 끝까지 갔으면 
		if isCover(m) || r == len(s) {
			// 커버 되면 양편 인덱스 길이를 계산해서 min 보다 작으면 리턴 값으로 기억
			if isCover(m) {
				if min_len == -1 || r-l < min_len {
					min_len = r-l 
					left, right = l, r
				}
			}

			// 왼편 인덱스의 현재 가리키는 문자를 맵에서 증가 시킴
			if _, exist := m[s[l]]; exist {
				m[s[l]] ++
			}

			l++

		} else {
			// 커버 되지 않은 상태이면 오른쪽 인덱스의 문자열을 맵에서 감소 시킴
			if _, exist := m[s[r]]; exist {
				m[s[r]] --
			}
			r++
		}

	}

	return s[left:right]
}

