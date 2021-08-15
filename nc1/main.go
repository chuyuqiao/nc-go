package main

import (
	"fmt"
	"strconv"
)

// 大数加法
func main() {
	s, t := "733064366", "459309139"
	// write code here
	s1, t1 := len(s), len(t)
	n := s1
	if t1 < s1 {
		n = t1
	}

	plus1 := 0
	var ret []string

	for i := 0; i < n; i++ {
		u, _ := strconv.Atoi(string(s[s1-1]))
		u2, _ := strconv.Atoi(string(t[t1-1]))
		sum := u + u2 + plus1
		if sum >= 10 {
			plus1 = 1
		} else {
			plus1 = 0
		}
		ret = append(ret, strconv.Itoa(sum%10))
		s1--
		t1--
	}

	if s1 > 0 {
		for i := 0; i < s1; i++ {
			u, _ := strconv.Atoi(string(s[s1-1]))
			sum := u + plus1
			if sum >= 10 {
				plus1 = 1
			} else {
				plus1 = 0
			}
			ret = append(ret, strconv.Itoa(sum%10))
		}
	}

	if t1 > 0 {
		for i := 0; i < t1; i++ {
			u2, _ := strconv.Atoi(string(t[t1-1]))
			sum := u2 + plus1
			if sum >= 10 {
				plus1 = 1
			} else {
				plus1 = 0
			}
			ret = append(ret, strconv.Itoa(sum%10))
		}
	}

	if plus1 > 0 {
		ret = append(ret, strconv.Itoa(plus1))
	}

	n = len(ret)
	for i := 0; i < n/2; i++ {
		ret[i], ret[n-1-i] = ret[n-1-i], ret[i]
	}

	rets := ""
	for i := 0; i < n; i++ {
		rets += ret[i]
	}
	fmt.Println(rets)
}
