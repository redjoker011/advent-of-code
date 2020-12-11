package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func problemOne(collection []string) int {
	validPword := 0
	for _, v := range collection {
		splitted := strings.Split(v, " ")
		var values = make([]string, 3, 3)
		for idx, s := range splitted {
			values[idx] = s
		}
		markers := strings.Split(values[0], "-")
		char := strings.Split(values[1], ":")

		match := strings.Count(values[2], char[0])

		min, _ := strconv.Atoi(markers[0])
		max, _ := strconv.Atoi(markers[1])
		if match >= min && match <= max {
			validPword++
		}
	}
	return validPword
}

func problemTwo(collection []string) int {
	validPword := 0
	for _, v := range collection {
		splitted := strings.Split(v, " ")
		var values = make([]string, 3, 3)
		for idx, s := range splitted {
			values[idx] = s
		}
		markers := strings.Split(values[0], "-")
		char := strings.Split(values[1], ":")

		var stringCollection []string
		for _, c := range values[2] {
			stringCollection = append(stringCollection, string(c))
		}

		min, _ := strconv.Atoi(markers[0])
		max, _ := strconv.Atoi(markers[1])

		nMin := min - 1
		nMax := max - 1

		// fmt.Printf("Collection %v Markers %v \n", stringCollection, markers)
		matchingChar := char[0]
		// fmt.Printf("Matcher %v-%v-%v \n", matchingChar, stringCollection[nMin], stringCollection[nMax])
		// fmt.Printf("Has Matching %v\n", stringCollection[nMin] == matchingChar || stringCollection[nMax] == matchingChar)
		// fmt.Printf("Has No Matching %v\n", !(stringCollection[nMin] == matchingChar && stringCollection[nMax] == matchingChar))
		if (stringCollection[nMin] == matchingChar) != (stringCollection[nMax] == matchingChar) {
			validPword++
		}
	}
	return validPword
}

func main() {
	data, err := os.Open("./input.txt")

	checkError(err)

	scanner := bufio.NewScanner(data)
	// scanner.Split(bufio.ScanWords)

	var collection []string

	for scanner.Scan() {
		content := scanner.Text()

		collection = append(collection, content)
	}

	ans := problemOne(collection)
	fmt.Printf("Problem 1 Answer: %v \n", ans)
	ans = problemTwo(collection)
	fmt.Printf("Problem 2 Answer: %v \n", ans)
}
