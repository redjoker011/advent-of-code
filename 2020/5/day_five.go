package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var rowsCount int = 127
var colsCount int = 7

var backMarker = "B"
var frontMarker = "F"
var leftMarker = "L"
var rightMarker = "R"

var lowerPointer = "0"
var higherPointer = "1"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func decConverter(char string) string {
	if char == frontMarker || char == leftMarker {
		return lowerPointer
	}
	return higherPointer
}

// Convert Code into Binary and then Convert Binary to Decimal using
// strconv package
func highestSeatID(passes []string) int64 {
	var best int64 = 0
	for i := 0; i < len(passes)-1; i++ {
		code := passes[i]
		rowCode := code[:7]
		colCode := code[7:]

		rowCode = strings.ReplaceAll(string(rowCode), frontMarker, lowerPointer)
		rowCode = strings.ReplaceAll(string(rowCode), backMarker, higherPointer)

		colCode = strings.ReplaceAll(string(colCode), leftMarker, lowerPointer)
		colCode = strings.ReplaceAll(string(colCode), rightMarker, higherPointer)

		row, _ := strconv.ParseInt(rowCode, 2, 64)
		col, _ := strconv.ParseInt(colCode, 2, 64)

		fmt.Printf("Row set to %v and Col set to %v\n", row, col)
		prod := (row * 8) + col
		best = Max(best, prod)
	}

	return best
}

func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	checkError(err)

	passes := strings.Split(string(data), "\n")

	fmt.Printf("Highest Seat ID %v", highestSeatID(passes))
}
