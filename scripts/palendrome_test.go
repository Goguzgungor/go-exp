package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"tek basamaklı sayı", 7, true},
		{"sıfır", 0, true},
		{"çift basamaklı palindrom", 11, true},
		{"çift basamaklı palindrom değil", 12, false},
		{"tek basamaklı uzun palindrom", 12321, true},
		{"çift basamaklı uzun palindrom", 123321, true},
		{"palindrom değil", 12345, false},
		{"leetcode örneği 121", 121, true},
		{"leetcode örneği 10", 10, false},
		{"negatif sayı", -121, false},
		{"negatif tek basamak", -1, false},
		{"sonu sıfır ile biten", 100, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.x); got != tt.want {
				t.Errorf("isPalindrome(%d) = %v, beklenen %v", tt.x, got, tt.want)
			}
		})
	}
}
