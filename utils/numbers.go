package utils

import "strconv"

func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func ParseNumber(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		Check(err)
		return -1
	}
	return val
}
