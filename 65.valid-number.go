/*
 * @lc app=leetcode id=65 lang=golang
 *
 * [65] Valid Number
 *
 * https://leetcode.com/problems/valid-number/description/
 *
 * algorithms
 * Hard (14.14%)
 * Likes:    451
 * Dislikes: 3340
 * Total Accepted:    128.5K
 * Total Submissions: 908.7K
 * Testcase Example:  '"0"'
 *
 * Validate if a given string can be interpreted as a decimal number.
 * 
 * Some examples:
 * "0" => true
 * " 0.1 " => true
 * "abc" => false
 * "1 a" => false
 * "2e10" => true
 * " -90e3   " => true
 * " 1e" => false
 * "e3" => false
 * " 6e-1" => true
 * " 99e2.5 " => false
 * "53.5e93" => true
 * " --6 " => false
 * "-+3" => false
 * "95a54e53" => false
 * 
 * Note: It is intended for the problem statement to be ambiguous. You should
 * gather all requirements up front before implementing one. However, here is a
 * list of characters that can be in a valid decimal number:
 * 
 * 
 * Numbers 0-9
 * Exponent - "e"
 * Positive/negative sign - "+"/"-"
 * Decimal point - "."
 * 
 * 
 * Of course, the context of these characters also matters in the input.
 * 
 * Update (2015-02-10):
 * The signature of the C++ function had been updated. If you still see your
 * function signature accepts a const char * argument, please click the reload
 * button to reset your code definition.
 * 
 */

func isDigit(c byte) bool  {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isNumber(s string) bool {
	ms := strings.TrimSpace(s)

	i := 0
	l := len(ms)

	ret := false

	// 부호이면 한칸 진행 
	if i < l && (ms[i] == '+' || ms[i] == '-') {
		i++
	}

	// 숫자일때까지 진행. 일단 valid 표시
	for i< l && isDigit(ms[i]) {
		i++
		ret = true
	}

	// 소숫점을 만나면 한칸 진행하고 내부에서 숫자를 추가로 진행 
	if i<l && ms[i] == '.' {
		i++
		// ret = false // "3." 케이스 true 로 리턴하기 위해 // 검증 데이터 오류로 보임
		for i<l && isDigit(ms[i]) {
			i++
			ret = true
		}
	}

	// 여기까지 ret 가 true 여야 e 를 체크할 수 있으므로 
	if ret == true && i<l && ms[i] == 'e' {
		i++
		ret = false

		// e이후에 부호 한칸 진행
		if i < l && (ms[i] == '+' || ms[i] == '-') {
			i++
		}
		// 숫자 진행
		for i< l && isDigit(ms[i]) {
			i++
			ret = true
		}
	}


	// 모든 체크가 끝나면 길이 끝까지 진행해야 함.
	// 추가 문자열이 남으면 잘못된 포맷임
	if i < l {
		ret = false
	} 

	return ret
}

