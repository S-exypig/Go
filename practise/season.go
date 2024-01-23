package practise

func Season(month int) string {
	switch month {
	case 1, 2, 3:
		return "Spring"
	case 4, 5, 6:
		return "Summer"
	case 7, 8, 9:
		return "Autumn"
	case 10, 11, 12:
		return "Winter"
	default:
		return "wfy"
	}
}
