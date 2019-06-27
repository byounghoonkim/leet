import "fmt"

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */
func multiply(num1 string, num2 string) string {
	ret := []byte{'0'}
	sc1 := 0
	sc2 := 0
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			r := mul(num1[i], num2[j])
			if r[0] != '0' {
				r = shift(r, sc1+sc2)
				fmt.Println("R", string(r))
				ret = add(ret, r)
				fmt.Println("OK", string(ret), string(r))
			}
			sc2++
		}
		sc2 = 0
		sc1++
	}
	return string(ret)
}

func mul(a, b byte) []byte {
	r := (a - '0') * (b - '0')

	if r/10 > 0 {
		return ([]byte{(r / 10) + '0', (r % 10) + '0'})
	} else {
		return ([]byte{(r % 10) + '0'})
	}
}

func shift(r []byte, s int) []byte {
	ret := r
	for i := 0; i < s; i++ {
		ret = append(ret, '0')
	}
	return ret
}

func add(a, b []byte) []byte {
	ret := []byte{}
	adder := []byte{}

	if len(a) > len(b) {
		ret = make([]byte, len(a))
		copy(ret, a)
		adder = b
	} else {
		ret = make([]byte, len(b))
		copy(ret, b)
		adder = a
	}

	var carry byte
	for i := 0; i < len(adder); i++ {
		r := adder[len(adder)-1-i] + ret[len(ret)-1-i] - '0' - '0' + carry
		carry = r / 10
		if len(ret)-1-i-1 >= 0 {
			ret[len(ret)-1-i-1] += carry
			carry = 0
		}
		r = r%10 + '0'
		ret[len(ret)-1-i] = r
	}

	if carry == 1 {
		return append([]byte{'1'}, ret...)
	} else {
		return ret
	}
}

