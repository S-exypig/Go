package practise

func BubSort(slInt []int) {
	for i := 0; i < len(slInt)-1; i++ {
		for j := 0; j < len(slInt)-i-1; j++ {
			if slInt[j] > slInt[j+1] {
				slInt[j], slInt[j+1] = slInt[j+1], slInt[j]
			}
		}
	}
}
