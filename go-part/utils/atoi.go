package utils

import "strconv"

func Atoi(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return num
}
