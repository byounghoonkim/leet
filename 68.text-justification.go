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
func fullJustify(words []string, maxWidth int) []string {
	ret := []string{}

	wi := []int{}
	wl := []int{}
	w_len := 0 
	s_len := 0
	for i, word := range(words) {
		if w_len + len(word) + s_len + 1 > maxWidth {
			wi = append(wi, i)
			wl = append(wl, w_len)
			w_len = 0
			s_len = 0
		} 

		if w_len != 0 {
			s_len ++
		}
		w_len += len(word)
	}

	for _,i  := range(wi) {

	}

	return ret
}

