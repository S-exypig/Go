package practise

func MapFunc(f func(int) int, slInt []int) (res []int) {
	for _, v := range slInt {
		res = append(res, f(v))
	}
	return
}
