package practise

func Fibonacci(n int) (pos int, ret int) {
	if n <= 1 {
		ret = 1
	} else {
		_, v1 := Fibonacci(n - 1)
		_, v2 := Fibonacci(n - 2)
		ret = v1 + v2
	}
	pos = n
	return
}
