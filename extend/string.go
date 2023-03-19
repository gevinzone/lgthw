package extend

func ReverseString(s string) string {
	ss := []rune(s)
	left, right := 0, len(ss)-1
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		left++
		right--
	}
	return string(ss)
}
