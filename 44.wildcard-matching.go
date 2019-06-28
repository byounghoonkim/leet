/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 *
 * https://leetcode.com/problems/wildcard-matching/description/
 *
 * algorithms
 * Hard (22.95%)
 * Likes:    1088
 * Dislikes: 76
 * Total Accepted:    180.7K
 * Total Submissions: 787.4K
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement wildcard pattern
 * matching with support for '?' and '*'.
 *
 *
 * '?' Matches any single character.
 * '*' Matches any sequence of characters (including the empty sequence).
 *
 *
 * The matching should cover the entire input string (not partial).
 *
 * Note:
 *
 *
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like ? or *.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 *
 *
 * Example 2:
 *
 *
 * Input:
 * s = "aa"
 * p = "*"
 * Output: true
 * Explanation: '*' matches any sequence.
 *
 *
 * Example 3:
 *
 *
 * Input:
 * s = "cb"
 * p = "?a"
 * Output: false
 * Explanation: '?' matches 'c', but the second letter is 'a', which does not
 * match 'b'.
 *
 *
 * Example 4:
 *
 *
 * Input:
 * s = "adceb"
 * p = "*a*b"
 * Output: true
 * Explanation: The first '*' matches the empty sequence, while the second '*'
 * matches the substring "dce".
 *
 *
 * Example 5:
 *
 *
 * Input:
 * s = "acdcb"
 * p = "a*c?b"
 * Output: false
 *
 *
 */

func match(s []byte, p []byte, si int, pi int, c [][]bool) bool {

	// 재귀 함수 제일 앞 단에서 인덱스 페어에 대한 캐시를 확인한다.
	if c[si][pi] == true {
		return false
	}

	// 두 인덱스 모두 끝에 도달했으면 매치 성공이다.
	if len(s) == si && len(p) == pi {
		return true
	}

	// s는 남았는데 p 가 끝에 도달 했다면 매치 실패이다.
	// 캐시 한다.
	if len(s) > si && len(p) == pi {
		c[si][pi] = true
		return false
	}

	// s 가 끝에 도달 했는데 *이 남았다면 *을 건너 뛰고 진행해 본다.
	if len(s) == si {
		if p[pi] == '*' {
			return match(s, p, si, pi+1, c)
		}

		// * 이외에는 매칭 실패이다.
		c[si][pi] = true
		return false
	}

	// * 매칭이 시작 되면
	if p[pi] == '*' {
		// 중복 *을 제거 하고
		j := 0
		for i := pi; i < len(p); i++ {
			if p[i] == '*' {
				j = i
			} else {
				break
			}
		}

		// *이 s[si] 부터 끝까지 매치 되는 것을 확인 하기 위해
		// 루프를 돌면서 재귀를 호출한다.
		if true == match(s, p, si+1, j, c) {
			return true
		}

		if true == match(s, p, si, j+1, c) {
			return true
		}
		/*
			for i := si; i <= len(s); i++ {
				if true == match(s, p, i, j+1, c) {
					return true
				}
			}
		*/

		// 재귀로 매칭 못 했다면 실패 이다.
		c[si][pi] = true
		return false
	}

	// ? 는 한 문자만 매칭 된 것으로 하고 재귀 호출 한다.
	if p[pi] == '?' {
		return match(s, p, si+1, pi+1, c)
	}

	// 그 외 문자는 하나씩 비교하여 매칭 한다.
	if s[si] == p[pi] {
		return match(s, p, si+1, pi+1, c)
	}

	// 여기까지 매칭 하지 못했다면 매칭 실패이다.
	c[si][pi] = true
	return false

}

func isMatch(s string, p string) bool {

	// len(s) * len(p) 크기의 불 슬라이스를 생성 해
	// 해당 인덱스가 검색에 실패 할때 true로 셋팅해
	// 동적 프로그래밍으로 접근한다.
	// 이렇게 하는 이유는 동일한 탐색을 반복 하는 과정에서
	// 타임 리미트 초과 가 발생하게 때문이다.
	cache := make([][]bool, len(s)+1)
	for i, _ := range cache {
		cache[i] = make([]bool, len(p)+1)
	}

	return match([]byte(s), []byte(p), 0, 0, cache)

	// 이 알고리즘은 정확히 복잡도 계산을 할 수가 없다.
	// TODO 이 알고리즘의 복잡도를 계산하는 과정을 블로그로 쓰자!!!
}

