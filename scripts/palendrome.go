package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	j := len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[i] != s[j] {
			return false
		} else {
			j--
		}
		if i == j {
			return true
		}
	}
	return true
}
