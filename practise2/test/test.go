package test

func Test(str string) string {
	if str == "hello" {
		str += " world"
	}
	return str
}
