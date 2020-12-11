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

	validPword := 0
	var matches []string

	for _, v := range collection {
		splitted := strings.Split(v, " ")
		var values = make([]string, 3, 3)
		for idx, s := range splitted {
			values[idx] = s
		}
		markers := strings.Split(values[0], "-")
		char := strings.Split(values[1], ":")
		fmt.Printf("Markers %v-%v \n", markers[0], markers[1])
		fmt.Printf("Char %s\n", char[0])

		match := 0
		for _, c := range values[2] {
			fmt.Printf("Char %s-%s\n", char[0], string(c))
			if char[0] == string(c) {
				fmt.Println("Match \n")
				match++
			}
		}

		min, _ := strconv.Atoi(markers[0])
		max, _ := strconv.Atoi(markers[1])
		if match >= min && match <= max {
			matches = append(matches, v)
			validPword++
		}
	}

	fmt.Printf("Matches: %v \n", matches)
	fmt.Printf("Answer: %v \n", validPword)
}
