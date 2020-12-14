package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
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
	validPasswords := 0
	for _, line := range lines {
		extracted := getParams(line)
		if testPasswordv2(extracted) {
			validPasswords++

		}
	}
	fmt.Println(validPasswords)
}

func testPassword(Params map[string]string) bool {
	count := 0
	for _, char := range Params["password"] {
		if string(char) == Params["letter"] {
			count++
		}
	}
	minChar, _ := strconv.Atoi(Params["minChar"])
	maxChar, _ := strconv.Atoi(Params["maxChar"])
	if minChar <= count {
		if count <= maxChar {
			return true
		}
	}
	return false
}

func testPasswordv2(Params map[string]string) bool {
	matched := false
	pos1, _ := strconv.Atoi(Params["minChar"])
	pos2, _ := strconv.Atoi(Params["maxChar"])

	for i, char := range Params["password"] {
		if string(char) == Params["letter"] {
			if i+1 == pos1 {
				matched = true
			} else if i+1 == pos2 && !matched {
				matched = true
			} else if i+1 == pos2 && matched {
				return false
			}
		}
	}

	if matched {
		fmt.Println(Params)
		return true
	}
	return false
}

func getParams(value string) (paramsMap map[string]string) {
	var regEx = `(?P<minChar>[\d]*)\-(?P<maxChar>[\d]*)\s*(?P<letter>[\w]*)\:\s(?P<password>[\w]*)`
	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(value)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}
