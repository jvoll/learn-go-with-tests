package iteration

func Repeat(input string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += input
	}
	return repeated
}
