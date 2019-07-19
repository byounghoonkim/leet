/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 *
 * https://leetcode.com/problems/text-justification/description/
 *
 * algorithms
 * Hard (23.71%)
 * Likes:    383
 * Dislikes: 1057
 * Total Accepted:    100.3K
 * Total Submissions: 423.2K
 * Testcase Example:  '["This", "is", "an", "example", "of", "text", "justification."]\n16'
 *
 * Given an array of words and a width maxWidth, format the text such that each
 * line has exactly maxWidth characters and is fully (left and right)
 * justified.
 * 
 * You should pack your words in a greedy approach; that is, pack as many words
 * as you can in each line. Pad extra spaces ' ' when necessary so that each
 * line has exactly maxWidth characters.
 * 
 * Extra spaces between words should be distributed as evenly as possible. If
 * the number of spaces on a line do not divide evenly between words, the empty
 * slots on the left will be assigned more spaces than the slots on the right.
 * 
 * For the last line of text, it should be left justified and no extra space is
 * inserted between words.
 * 
 * Note:
 * 
 * 
 * A word is defined as a character sequence consisting of non-space characters
 * only.
 * Each word's length is guaranteed to be greater than 0 and not exceed
 * maxWidth.
 * The input array words contains at least one word.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * words = ["This", "is", "an", "example", "of", "text", "justification."]
 * maxWidth = 16
 * Output:
 * [
 * "This    is    an",
 * "example  of text",
 * "justification.  "
 * ]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * words = ["What","must","be","acknowledgment","shall","be"]
 * maxWidth = 16
 * Output:
 * [
 * "What   must   be",
 * "acknowledgment  ",
 * "shall be        "
 * ]
 * Explanation: Note that the last line is "shall be    " instead of "shall
 * be",
 * because the last line must be left-justified instead of fully-justified.
 * ⁠            Note that the second line is also left-justified becase it
 * contains only one word.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input:
 * words =
 * ["Science","is","what","we","understand","well","enough","to","explain",
 * "to","a","computer.","Art","is","everything","else","we","do"]
 * maxWidth = 20
 * Output:
 * [
 * "Science  is  what we",
 * ⁠ "understand      well",
 * "enough to explain to",
 * "a  computer.  Art is",
 * "everything  else  we",
 * "do                  "
 * ]
 * 
 * 
 */

 func spaces(n int) string {
	 str := ""
	 for i:= 0 ; i < n ; i++ {
		 str += " "
	 }
	 return str
 }

 func justify(words []string, maxWidth int) string {

	if len(words) == 0 {
		return ""
	}

	if len(words) == 1 {
		return words[0] + spaces(maxWidth-len(words[0]))
	}

	wl := 0
	for _, w := range(words) {
		wl += len(w)
	}

	str := words[0] 
	for i := 1; i < len(words); i++ {
		
		// 각 단어를 돌면서 단어를 연결한다.
		// 연결 시 사이 공백의 갯수를 정해야 하는데
		// maxWidth 에서 단어 전체의 길이를 뺀 값이 사용할 공백 전체의 갯수가 된다.
		// 이 공백 전체의 갯수를 단어갯수-1 로 나누고 
		// 위 나누기의 나머지 값을 앞에서 부터 1씩 더해 채워 공백을 만들어야 한다.
		// 공백이 10개 이고 단어가 5개라고 하면
		// 공백은 3 3 2 2 처럼 추가 되어야 한다.

		sc := (maxWidth-wl) / (len(words)-1)
		if (maxWidth-wl) % (len(words)-1) > i - 1 {
			sc++
		}

		str += spaces(sc)
		str += words[i]
	}

	return str
 }

func fullJustify(words []string, maxWidth int) []string {
	ret := []string{}

	// 단어를 나눌 인덱스를 미리 구한다.
	// 구하는 방법은 각 단어에 공백을 하나 추가한 길이를 더해 가며 maxWidth를 초과 하기 전까지
	// 진행하고 초과했을 때의 인덱스를 저장하는 방식이다.
	wi := []int{}
	w_len := 0 
	for i, word := range(words) {
		if w_len + len(word) >= maxWidth {
			wi = append(wi, i)
			w_len = 0
		} 

		// 0 이면 첫 단어라 공백 길이를 추가하지 않는다.
		if w_len > 0 {
			w_len ++
		}
		w_len += len(word) 
	}

	// 인덱스 슬라이스를 돌면서
	// 구간 별 서브 워드 슬라이스를 만들어 justify 함수를 호출한다.
	for i, v := range(wi) {
		s, e := 0, 0
		if i == 0 {
			s = 0
		} else {
			s = wi[i-1]
		}
		e = v
		sub := words[s:e]
		if len(sub) > 0 {
			str := justify(sub, maxWidth)
			ret = append(ret, str)
		}
	}

	// 마지막 구간은 위에서 처리 되지 않으므로 따로 구해서 처리한다.
	sub := words
	if len(wi) > 0 {
		sub = words[wi[len(wi)-1]:]
	} 

	if len(sub) > 0 {
		str := sub[0]
		for i:= 1 ; i < len(sub); i++ {
			str += " "
			str += sub[i]
		}

		// 마지막 구간은 단어가 모자라(단어가 하나거나 등) 공백을 중간에 채워넣지 못할 수도 있다.
		// 이때는 마지막에 공백을 다 채워서 justify 해 주어야 검증 시 pass 된다.
		str += spaces(maxWidth-len(str))

		ret = append(ret, str)
	}

	return ret
}

