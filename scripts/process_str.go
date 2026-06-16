package main

func removeLast(s string) string {
	if s == "" {
		return s
	}
	return s[:len(s)-1]
}

func dublicate(s string) string {
	return s + s
}

func reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func processStr(s string) string {

	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			result = removeLast(result)
		} else if s[i] == '#' {
			result = dublicate(result)
		} else if s[i] == '%' {
			result = reverse(result)
		} else {
			result += string(s[i])
		}
	}
	return result
}
