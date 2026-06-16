package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{"leetcode örneği 1", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"nums2 boş", []int{1}, 1, []int{}, 0, []int{1}},
		{"nums1 boş", []int{0}, 0, []int{1}, 1, []int{1}},
		{"araya geçen değerler", []int{2, 4, 6, 0, 0, 0}, 3, []int{1, 3, 5}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"nums2 tamamen daha büyük", []int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"nums2 tamamen daha küçük", []int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"tekrar eden değerler", []int{1, 1, 0, 0}, 2, []int{1, 1}, 2, []int{1, 1, 1, 1}},
		{"negatif sayılar", []int{-3, 0, 0, 0}, 1, []int{-2, -1, 5}, 3, []int{-3, -2, -1, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.want) {
				t.Errorf("merge() = %v, beklenen %v", tt.nums1, tt.want)
			}
		})
	}
}
