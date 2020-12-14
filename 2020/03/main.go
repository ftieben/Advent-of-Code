package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Print(loadInput())
}

//Todo: there must be a better way then to define the array so early
func loadInput() [323][31]string {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	var Map [323][31]string
	lines := strings.Split(string(dat), "\n")
	for i, row := range lines {
		for n, column := range row {
			Map[i][n] = string(column)
		}
	}

	return Map
}
