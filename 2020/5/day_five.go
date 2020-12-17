package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
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
func highestSeatID(passes []string) (hSeat int64, seatId int) {
	var best int64 = 0
	var ids []int
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

		prod := (row * 8) + col
		ids = append(ids, int(prod))
		best = Max(best, prod)
	}

	// Get Seat ID Where preceding seat is - 1 and next seat is + 1
	// e.g 1,2,3 where 2 is id
	var lastId int = -1
	var matchSeatId int = 0
	sort.Ints(ids) // always sort when finding linear integers
	for _, i := range ids {
		if lastId != -1 && i-lastId == 2 {
			matchSeatId = i - 1
		}
		lastId = i
	}
	return best, matchSeatId
}

func seatSearch(passes []string) (highesSeat, ExactSeat int) {
	var hSeat, seatId int
	var ids []int

	for c := 0; c < len(passes)-1; c++ {
		code := passes[c]

		var row float64 = 0
		var col float64 = 0

		for i := 0; i < 7; i++ {
			if string(code[i]) == "B" {
				val := float64(6 - i)
				row += math.Pow(float64(2), val)
			}
		}

		for i := 7; i < 10; i++ {
			if string(code[i]) == "R" {
				val := 9 - i
				ans := math.Pow(float64(2), float64(val))
				col += ans
			}
		}

		rowCol := int(row)*8 + int(col)
		ids = append(ids, int(rowCol))

		if hSeat < rowCol {
			hSeat = int(rowCol)
		}
	}

	var lastId = -1
	sort.Ints(ids)
	// Find the missing seat in list of sorted seat ids
	// i.e 556, 558 missing 557 in between
	for _, v := range ids {
		if lastId != -1 && v-lastId == 2 {
			seatId = v - 1
		}
		lastId = v
	}

	return hSeat, seatId
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
	hSeat, seatID := highestSeatID(passes)
	fmt.Printf("Highest Seat ID %v \n", hSeat)
	fmt.Printf("Seat ID %v\n", seatID)

	// Using Binary Conversion
	highest, exact := seatSearch(passes)
	fmt.Printf("Highest Seat ID %v and Exact Seat ID %v \n", highest, exact)
}
