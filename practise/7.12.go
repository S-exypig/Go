package practise

func SplitStr(oriStr string, i int) (str1, str2 string) {
	str1, str2 = oriStr[:i], oriStr[i+1:]
	return
}
