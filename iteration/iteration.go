package iteration

import "strings"

func Repeat(str string, count int) string {
	if count <= 0 {
		return str
	}
	var repeated string
	for i := 0; i < count; i++ {
		repeated += str
	}

	return repeated
}

func ImprovedRepeat(s string, count int) string {
	var repeated strings.Builder

	for i := 0; i < count; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
