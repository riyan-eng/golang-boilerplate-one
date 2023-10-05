package util

import "strconv"

func StringNumToInt(strnum string) (num int) {
	num, err:= strconv.Atoi(strnum)
	PanicIfNeeded(err)
	return
}