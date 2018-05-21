package iteration

// Repeat repeats the passed in string...
func Repeat(s string, count int) string {
	var r string
	for i := 0; i < count; i++ {
		r += s
	}
	return r
}
