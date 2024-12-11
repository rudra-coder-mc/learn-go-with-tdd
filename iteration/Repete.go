package iteration

func Repeat(input string, times int) string {
	out := ""
	for i := 0; i < times; i++ {
		out += input
	}
	return out
}
