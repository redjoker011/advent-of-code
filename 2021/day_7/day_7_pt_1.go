package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2021/day_7/input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")

	var pos []int
	// Iterate Over Array and find Max Position
	max := 0
	for _, s := range strings.Split(input[0], ",") {
		i, err := strconv.Atoi(s)

		if err != nil {
			log.Fatal(err)
		}

		if i > max {
			max = i
		}

		pos = append(pos, i)
	}

	// Create a single dimension Array that has length equal to maximum Position
	var dists []int = make([]int, max+1)
	for i := 0; i <= 1; i++ {
		// Iterate over positions to check which move a certain is efficient
		for _, v := range pos {
			dists[i] += int(math.Abs(float64(v - i)))
			fmt.Printf("ABS %v; V %v; I %v \n", int(math.Abs(float64(v-i))), v, i)
		}
	}

	// Sort out the least move
	min := 999999999999
	for _, v := range dists {
		if v < min {
			min = v
		}
	}

	fmt.Println(min)
}
