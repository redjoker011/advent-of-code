package main

import (
	// "bufio"
	"fmt"
	"io/ioutil"
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

			if innerCounter == len(reqFields) {
				validPassportCount++
			}
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

	fmt.Printf("Answer to Problem 1 %v", FindValidPassports(passports))
}
