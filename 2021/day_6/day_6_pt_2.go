package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2021/day_6/input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")

	var fish = make([]int, 9)
	for _, i := range strings.Split(input[0], ",") {
		n, err := strconv.Atoi(i)

		if err != nil {
			log.Fatal(err)
		}
		fish[n]++
	}

	// Loop for 256 days
	for j := 0; j < 256; j++ {
		fmt.Printf("Day %v \n", j)
		fish = step(fish)
	}

	sum := 0
	for _, v := range fish {
		sum += v
	}

	fmt.Println("Sum %v", sum)
}

func step(fish []int) []int {
	var next = make([]int, 9)

	for i := 1; i < 9; i++ {
		next[i-1] = fish[i]
	}

	next[6] += fish[0]
	next[8] += fish[0]
	return next
}
