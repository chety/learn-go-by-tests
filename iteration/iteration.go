package iteration

func Repeat(str string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += str
	}
	if len(repeated) == 0 {
		return str
	}

	return repeated
}
