package iteration

func Repeat(s string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated = repeated + s
	}
	return repeated

}
