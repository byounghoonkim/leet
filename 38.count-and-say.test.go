package main

import "strconv"
import "fmt"

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

func say(s []byte) []byte {
	// 이 함수는 입력된 문자열에 count and say를 적용한 문자열을 리턴한다.
	ret := make([]byte, 0, len(s)*2)

	cur := s[0]
	count := 0
	for _, c := range s {
		if cur != c {
			// 이전 문자와 다른 문자가 나오면 이때까지 카운트 한 내용을 리턴 문자열에 추가한다.
			if count > 0 {
				ret = append(ret, []byte(strconv.Itoa(count))...)
				ret = append(ret, cur)

				// 그리고 현재 문자열을 초기화 한다.
				cur = c
				count = 1
			}
		} else {
			count += 1
		}
	}

	// 마지막에 카운트한 내용도 반영해야 한다.
	if count > 0 {
		ret = append(ret, []byte(strconv.Itoa(count))...)
		ret = append(ret, cur)
	}

	return ret
}

func countAndSay(n int) string {

	cur := []byte{'1'}

	for i := 1; i < n; i++ {
		cur = say(cur)
	}

	return string(cur)
}

func main() {
	fmt.Println(countAndSay(50))
}
