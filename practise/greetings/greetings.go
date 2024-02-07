package greetings

import "time"

func ISAM() bool {
	if time.Now().Hour() < 12 {
		return true
	}
	return false
}

func IsAfternoon() bool {
	if hour := time.Now().Hour(); hour >= 12 && hour < 18 {
		return true
	}
	return false
}

func IsEvening() bool {
	if time.Now().Hour() >= 18 {
		return true
	}
	return false
}
