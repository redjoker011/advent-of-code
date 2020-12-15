package main

import (
	// "bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"

	// "os"
	"strings"
)

var reqFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func FindValidPassports(passports []string) int {
	validPassportCount := 0

	for _, v := range passports {
		innerCounter := 0
		newText := strings.Join(strings.Split(v, "\n"), " ")

		for _, f := range reqFields {
			if strings.Contains(newText, f) {
				innerCounter++
			}
		}

		if innerCounter == len(reqFields) {
			validPassportCount++
		}
	}
	return validPassportCount
}

func FindValidPassportsStrict(passports []string) int {
	validPassportCount := 0

	for _, v := range passports {
		innerCounter := 0
		newText := strings.Split(strings.Join(strings.Split(v, "\n"), " "), " ")

		details := make(map[string]string)
		for _, v := range newText {
			values := strings.Split(v, ":")
			details[values[0]] = values[1]
		}

		for _, f := range reqFields {
			detail := details[f]

			// fmt.Printf("Matchers %v %v %v \n", detail, details, f)
			switch f {
			case "byr":
				bYear, _ := strconv.Atoi(detail)
				if bYear >= 1920 && bYear <= 2002 {
					innerCounter++
				}
			case "iyr":
				iYear, _ := strconv.Atoi(detail)
				if iYear >= 2010 && iYear <= 2020 {
					innerCounter++
				}
			case "eyr":
				expY, _ := strconv.Atoi(detail)
				if expY >= 2020 && expY <= 2030 {
					innerCounter++
				}
			case "hgt":
				re := regexp.MustCompile(`[0-9]{2}in`)
				if re.MatchString(detail) {
					measurement, _ := strconv.Atoi(detail[:2])
					if measurement >= 59 && measurement <= 76 {
						innerCounter++
					}
				}
				re = regexp.MustCompile(`[0-9]{3}cm`)
				if re.MatchString(detail) {
					measurement, _ := strconv.Atoi(detail[:3])
					if measurement >= 150 && measurement <= 193 {
						innerCounter++
					}
				}
			case "hcl":
				re := regexp.MustCompile(`#([a-f0-9]{6})`)
				if re.MatchString(detail) {
					innerCounter++
				}
			case "ecl":
				colorCodes := "amb blu brn gry grn hzl oth"
				if strings.Contains(colorCodes, detail) {
					innerCounter++
				}
			case "pid":
				re := regexp.MustCompile(`[0-9]{9}`)
				if re.MatchString(detail) {
					innerCounter++
				}
			}
		}

		// fmt.Printf("Inner Counter %v, %v\n", innerCounter, len(reqFields))
		if innerCounter == len(reqFields) {
			// fmt.Printf("Match %v\n", details)
			validPassportCount++
		}
	}
	return validPassportCount
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	checkError(err)
	var passports []string
	content := strings.Split(string(data), "\n\n")

	for i := 0; i < (len(content) - 1); i++ {
		passports = append(passports, content[i])
	}

	fmt.Printf("Answer to Problem 1 %v\n", FindValidPassports(passports))
	fmt.Printf("Answer to Problem 2 %v\n", FindValidPassportsStrict(passports))
}
