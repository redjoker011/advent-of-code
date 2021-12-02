package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2021/day_2/input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")
	var depth int
	var horizontal int

	for idx, v := range input {

		if idx == len(input)-1 {
			fmt.Println("Reached EOF \n")
			break
		}

		splittedValue := strings.Split(v, " ")
		move := splittedValue[0]
		count, _ := strconv.Atoi(splittedValue[1])

		switch move {
		case "up":
			depth -= count
		case "down":
			depth += count
		default:
			horizontal += count
		}
	}

	fmt.Printf("Product of depth: %v and horizontal: %v is %v", depth, horizontal, depth*horizontal)
}
