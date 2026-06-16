package main

import "testing"

func TestProcessStr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"özel karakter yok", "abc", "abc"},
		{"boş girdi", "", ""},
		{"yıldız son karakteri siler", "abc*", "ab"},
		{"yıldız ortada", "a*b", "b"},
		{"boş sonuçta yıldız güvenli", "*", ""},
		{"diyez sonucu ikiler", "ab#", "abab"},
		{"çift diyez", "a##", "aaaa"},
		{"yüzde sonucu ters çevirir", "abc%", "cba"},
		{"diyez sonra yüzde", "ab#%", "baba"},
		{"yıldız sonra diyez", "ab*#", "aa"},
		{"yüzde sonra yıldız", "ab%*", "b"},
		{"birden fazla operasyon", "abc*%#", "baba"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processStr(tt.s); got != tt.want {
				t.Errorf("processStr(%q) = %q, beklenen %q", tt.s, got, tt.want)
			}
		})
	}
}
