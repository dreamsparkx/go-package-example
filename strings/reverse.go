package strings

// Reverse asd
func Reverse(s string) string {
	runes := []rune(s)
	reversedRunes := reverseRunes(runes)
	return string(reversedRunes)
}
