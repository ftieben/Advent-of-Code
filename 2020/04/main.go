package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type passport struct {
	byr int    // (Birth Year)
	iyr int    // (Issue Year)
	eyr int    // (Expiration Year)
	hgt string // (Height)
	hcl string // (Hair Color) // hexcode
	ecl string // (Eye Color) // string
	pid string // (Passport ID)
	cid int    // (Country ID)
}

type passportRule struct { // true to ignore
	byr bool // (Birth Year)
	iyr bool // (Issue Year)
	eyr bool // (Expiration Year)
	hgt bool // (Height)
	hcl bool // (Hair Color) // hexcode
	ecl bool // (Eye Color) // string
	pid bool // (Passport ID)
	cid bool // (Country ID)
}

func main() {
	passports := loadInput("input.txt")
	var rule passportRule
	rule.cid = true
	validPassports := 0
	for _, p := range passports {
		if isPassportValid(p, rule) {
			validPassports++
		}
	}
	fmt.Println("-----------------------------------------------------------\n",
		validPassports,
		"from",
		len(passports),
		"Passports are valid \n-----------------------------------------------------------")
}

func inRange(value int, from int, to int) bool {
	if from < value {
		if value < to {
			return true
		}
	}
	return false
}

func hasNdigits(value int, n int) bool {
	if len(strconv.Itoa(value)) == n {
		return true
	}
	return false
}

func isDigit(str string) bool {
	out := true
	for _, v := range str {
		out = out && unicode.IsDigit(v)
	}
	return out
}

func hair(str string) bool {
	rest := str[:1]
	if str[0] != '#' || len(rest) == 6 {
		return false
	}
	_, err := strconv.ParseInt("0x"+rest, 0, 64)
	return err == nil
}

func isPassportValid(passport passport, rule passportRule) bool {

	if !rule.byr {
		//byr (Birth Year) - four digits; at least 1920 and at most 2002.
		if passport.byr == 0 || !inRange(passport.byr, 1920, 2002) || !hasNdigits(passport.byr, 4) {
			return false
		}
	}

	if !rule.iyr {
		//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		if passport.iyr == 0 || !inRange(passport.iyr, 2010, 2020) || !hasNdigits(passport.iyr, 4) {
			return false
		}
	}

	if !rule.eyr {
		//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		if passport.eyr == 0 || !inRange(passport.eyr, 2020, 2030) || !hasNdigits(passport.eyr, 4) {
			return false
		}
	}

	if !rule.hgt {
		//hgt (Height) - a number followed by either cm or in:
		//		If cm, the number must be at least 150 and at most 193.
		//		If in, the number must be at least 59 and at most 76.
		if passport.hgt == "" {
			return false
		}
	}

	if !rule.hcl {
		//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		if passport.hcl == "" {
			return false
		}
	}

	if !rule.ecl {
		//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		if passport.ecl == "" {
			return false
		}
	}

	if !rule.pid {
		//pid (Passport ID) - a nine-digit number, including leading zeroes.
		if passport.pid == "" || !isDigit(passport.pid) {
			return false
		}
		pidDigits, _ := strconv.Atoi(passport.pid)
		if !hasNdigits(pidDigits, 9) {
			return false
		}

	}

	if !rule.cid {
		if passport.cid == 0 {
			return false
		}
	}

	return true

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadInput(filename string) []passport {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	lines := strings.Split(string(dat), "\n")
	var passports []passport

	tmpString := ""
	for _, row := range lines {
		if row != "" && tmpString == "" {
			tmpString = row
		} else if row != "" {
			tmpString = tmpString + " " + row
		} else {
			passports = append(passports, convertParams(getParams(tmpString)))
			tmpString = ""
		}
	}
	return passports
}

func emptyStringIsZero(emptyone string) string {
	if emptyone == "" {
		return "0"
	}
	return emptyone
}

func convertParams(input map[string]string) passport {
	var passport passport
	var err error
	fmt.Println(input)
	passport.byr, err = strconv.Atoi(emptyStringIsZero(input["byr"]))
	check(err)
	passport.iyr, err = strconv.Atoi(emptyStringIsZero(input["iyr"]))
	check(err)
	passport.eyr, err = strconv.Atoi(emptyStringIsZero(input["eyr"]))
	check(err)
	passport.hgt = input["hgt"]
	passport.hcl = input["hcl"]
	passport.ecl = input["ecl"]
	passport.pid = input["pid"]
	check(err)
	passport.cid, err = strconv.Atoi(emptyStringIsZero(input["cid"]))
	check(err)
	return passport
}

func getParams(value string) (paramsMap map[string]string) {
	re := regexp.MustCompile("(ecl|hgt|cid|byr|eyr|iyr|pid|hcl):(#*[a-z0-9]+)")
	result := re.FindAllStringSubmatchIndex(value, -1)
	paramsMap = make(map[string]string)
	for _, match := range result {
		key := value[match[2]:match[3]]
		value := value[match[4]:match[5]]
		paramsMap[key] = value
	}
	return paramsMap
}
