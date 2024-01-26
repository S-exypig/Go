package practise

func TravelRune(arr []rune) (res []rune) {
	var lastRune rune
	for _, r := range arr {
		if r != lastRune {
			res = append(res, r)
		}
		lastRune = r
	}
	return
}
