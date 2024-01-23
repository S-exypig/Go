package practise

import "unicode/utf8"

func CountByteAndRune(s string) (byteCount int, runeCount int) {
	byteCount = len(s)
	runeCount = utf8.RuneCountInString(s)
	return byteCount, runeCount
}
