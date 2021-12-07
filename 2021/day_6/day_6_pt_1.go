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

	var lanternFish []int

	for _, i := range strings.Split(input[0], ",") {
		n, err := strconv.Atoi(i)

		if err != nil {
			log.Fatal(err)
		}
		lanternFish = append(lanternFish, n)
	}

	// Loop for 80 days
	for j := 0; j < 80; j++ {
		for i, l := range lanternFish {
			// Check if internal timer is zero
			// - Refresh timer to 6
			// - Spawn new lantern fish with timer of 8
			if l == 0 {
				lanternFish[i] = 6
				lanternFish = append(lanternFish, 8)
			} else {
				// Decrement counter
				lanternFish[i]--
			}
		}
	}

	fmt.Println(len(lanternFish))
}
