package practise

func FibonacciClosure() func() int {
	v1, v2 := 1, 1
	return func() int {
		v1, v2 = v2, v1+v2
		return v2
	}
}
