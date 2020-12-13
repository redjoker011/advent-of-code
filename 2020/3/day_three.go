package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var tree string = "#"

func problemOne(terrain []string, down int, right int) int {
	xOffset := 0
	yOffset := 0
	match := 0

	for yOffset < len(terrain) {
		line := terrain[yOffset]

		if line == "" {
			break
		}

		if string(line[xOffset]) == tree {
			match++
		}

		yOffset += down
		xOffset += right
		// Using Modulo
		// xOffset = (xOffset + right) % len(line)

		// Since modulo is slower we resolve into a much faster workaround
		if xOffset >= len(line) {
			xOffset -= len(line)
		}
	}
	return match
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	checkError(err)

	routeMap := strings.Split(string(data), "\n")
	fmt.Printf("Answer to Problem 1 is %v\n", problemOne(routeMap, 1, 3))

	fmt.Println("Answer to part 2:", (problemOne(routeMap, 1, 1) *
		problemOne(routeMap, 1, 3) *
		problemOne(routeMap, 1, 5) *
		problemOne(routeMap, 1, 7) *
		problemOne(routeMap, 2, 1)))
}
