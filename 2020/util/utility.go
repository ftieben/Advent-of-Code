package util

import (
	"strconv"
	"strings"
)

func intToIntArray(input int) []int {
	tmp := strconv.Itoa(input)
	tmpArrayOfLetters := strings.Split(tmp, "")

	var cups []int
	for _, value := range tmpArrayOfLetters {
		_tmp, err := strconv.Atoi(value)
		check(err)
		cups = append(cups, _tmp)
	}
	return cups
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
