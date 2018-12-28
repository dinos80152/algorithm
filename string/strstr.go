package string

func Strstr(a, b string) int {
	bLen := len(b)
	for i := 0; i <= len(a)-bLen; i++ {
		substr := a[i : i+bLen]
		for j := 0; j < bLen; j++ {
			if substr[j] != b[j] {
				break
			}
			if j == bLen-1 {
				return i
			}
		}
	}
	return -1
}
