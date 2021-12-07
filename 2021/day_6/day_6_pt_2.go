package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2021/day_6/input_t.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")

	// Breeding take 9 days so create a variable that will track the count of each
	// fish breeding timer
	var fish = make([]int, 9)
	for _, i := range strings.Split(input[0], ",") {
		n, err := strconv.Atoi(i)

		if err != nil {
			log.Fatal(err)
		}

		// Track fish based on timer 0 - 8 days
		fish[n]++
	}
	fmt.Println(fish)

	// Loop for 256 days
	for j := 0; j < 256; j++ {
		fmt.Printf("Day %v; Fish %v \n", j+1, fish)
		// Update Tracker each day
		fish = step(fish)
	}

	// Get the sum
	sum := 0
	for _, v := range fish {
		sum += v
	}

	fmt.Println("Sum %v", sum)
}

// Update Tracker
func step(fish []int) []int {
	// Create a variable to track breeding timer
	var next = make([]int, 9)

	// Reassign timer from previous day to current day
	// 0 = 1; 1 = 2 and so on..
	for i := 1; i < 9; i++ {
		next[i-1] = fish[i]
	}

	fmt.Printf("N %v \n", next)

	// Assume fish[0] will breed new fish reset timer
	// for new breed and parent
	next[6] += fish[0]
	next[8] += fish[0]

	fmt.Printf("M %v \n", next)
	return next
}
