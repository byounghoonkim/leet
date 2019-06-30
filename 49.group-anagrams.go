/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

func checkAnagrams(a, b string) bool {

	if len(a) != len(b) {
		return false
	}

	checker := make(map[rune]int)

	for _, r := range a {
		checker[r] += 1
	}

	for _, r := range b {
		checker[r] -= 1
	}

	for _, c := range checker {
		if c != 0 {
			return false
		}
	}

	return true
}

func myHash(w string) int {
	h := 0

	for i := 0; i < len(w); i++ {
		h += int(w[i])
	}

	return h
}

func addGroup(groupmap map[int][][]string, w string) {

	// 단어의 해시를 만든다.
	h := myHash(w)

	// 해시를 기반으로 맵에 접근해 값이 있다면 해당 그룹들을 순회 하며
	// 애너그램 그룹이 맞는지 최종 확인한다.
	if groups, ok := groupmap[h]; ok {

		for i, group := range groups {
			if true == checkAnagrams(group[0], w) {
				// 애너그램이 맞으면 해당 그룹에 단어를 추가한다.
				groups[i] = append(groups[i], w)
				return
			}
		}

		// 맞는 그룹이 없다면 해당 맵의 해시 아래 w 를 가진 새로운 그룹을 추가한다.
		groupmap[h] = append(groupmap[h], []string{w})

	} else {

		// 해시로 맵에 접근이 실패 했다면
		// 해당 맵으로 [][]string을 추가 하고 거기 w 를 가진 []string을 추가한다.
		groupmap[h] = [][]string{[]string{w}}
	}
}

func groupAnagrams(strs []string) [][]string {
	// int 를 키로 하는 []string을 배열로 가지는 맵을 만들어
	// 해시(맵) 을 기반으로 검색 성능을 높인다.
	// 이렇게 성능을 높이지 않으면 time limit exceed 에러가 발생한다.
	groupmap := make(map[int][][]string)

	// 각 단어를 그룹에 맞게 추가 한다.
	for _, w := range strs {
		addGroup(groupmap, w)
	}

	// 맵에서 그룹 부분만 추려 리턴 값을 만든다.
	ret := [][]string{}

	for _, groups := range groupmap {
		ret = append(ret, groups...)
	}

	return ret
}

