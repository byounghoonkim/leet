/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */

 func find_in_words(words []string, word_hit []int, word string) (found bool, index int) {
	 for i, a := range words {
		 if 0 == word_hit[i] {
			if 0 == strings.Compare(a, word) {
				return true, i
			}
		 }
	 }

	 return false, -1
 }

 func check_all_and_once_hit(word_hit []int) bool {
	 for _, a := range word_hit {
		 if a != 1 {
			 return false
		 }
	 }
	return true
 }

 func clear_word_hit(word_hit *[]int) {
	 for i, _ := range *word_hit {
		 (*word_hit)[i] = 0
	 }
 }

func findSubstring(s string, words []string) []int {
	var ret []int
	if 0 == len(words) {
		return ret
	}

	word_hit := make([]int, len(words))
	word_size := len(words[0])

	for i := 0 ; i < word_size ; i++ {
		for j := i ; j < len(s); j += word_size {
			if len(s) - j < len(words) * word_size {
				break
			}
			
			for k, _ := range words {
				word_start := j + k * word_size 
				word_end := word_start + word_size

				if word_end > len(s) {
					break
				}

				word := s[word_start:word_end]

				if found, index := find_in_words(words, word_hit, word) ; true == found {
					word_hit[index] += 1
				} else {
					break
				}
			}

			if true == check_all_and_once_hit(word_hit) {
				ret = append(ret, j)
			}
			clear_word_hit(&word_hit)
		}
	}

	return ret
}

