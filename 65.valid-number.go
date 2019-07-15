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

func isRealNumber(s string) bool {

	return false
}

func isNatureNumber(s string) bool {
	return false

}

func checkSpace(s string) bool{
	i:= 0
	for i = 0; i < len(s) ; i++ {
		if s[i] == ' ' {
			break
		}
	}

	for ; i < len(s) ; i++ {
		if s[i] != ' ' {
			break
		}
	}

	for ; i < len(s); i++ {
		if s[i] != ' ' {
			return false
		}
	}

	return true
}

func trimSpace(s string) string {
	ret := s
	i:= 0

	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}

	ret = ret[i:]

	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}

	ret = ret[:i]

	return ret

}

func isNumber(s string) bool {
	if false == checkSpace(s) {
		return false
	}

	fmt.Println(trimSpace(s))

	return false
    
}

