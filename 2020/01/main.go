package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	// Assign cap to avoid resize on every append.
	nums := make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		check(err)
		nums = append(nums, n)
	}
	testNumbers(nums)
	testNumbersv2(nums)
	//fmt.Print(string(dat))
}

func testNumbers(numbers []int) {
	finished := false
	for _, number1 := range numbers {
		for _, number2 := range numbers {
			summe := number1 + number2
			if summe == 2020 {
				fmt.Printf("number1=%d number2=%d result=%d\n", number1, number2, number1*number2)
				finished = true
			}
			if finished {
				break
			}
		}
		if finished {
			break
		}
	}
}

func testNumbersv2(numbers []int) {
	finished := false
	for _, number1 := range numbers {
		for _, number2 := range numbers {
			for _, number3 := range numbers {
				summe := number1 + number2 + number3
				if summe == 2020 {
					fmt.Printf("number1=%d number2=%d number3=%d result=%d\n", number1, number2, number3, number1*number2*number3)
					finished = true
				}
				if finished {
					break
				}
			}
		}
		if finished {
			break
		}
	}
}
