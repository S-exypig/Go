package practise

import "strings"

func StringsMap(str string) (res string) {
	res = strings.Map(func(r rune) rune {
		if r > 127 {
			return '?'
		}
		return r
	}, str)
	return
}
