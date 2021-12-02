package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2021/day_1/input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")
	initialDepth, _ := strconv.Atoi(input[0])

	counter := 0

	for idx, v := range input {
		depth, _ := strconv.Atoi(v)

		if idx == len(input)-1 {
			fmt.Println("Reached EOF")
			break
		} else if depth > initialDepth {
			counter++
		}

		initialDepth = depth
	}

	fmt.Printf("Depth Measurement %v", counter)
}
