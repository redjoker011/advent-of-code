package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var rowsCount int = 127
var colsCount int = 7

var backMarker = "B"
var frontMarker = "F"
var leftMarker = "L"
var rightMarker = "R"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func highestSeatID(passes []string) int {
	best := 0
	for i := 0; i < len(passes)-1; i++ {
		code := passes[i]
		rowCode := code[:7]
		colCode := code[7:]

		row := 0
		col := 0

		lPointer := 0
		rPointer := 0

		for r := 0; r <= len(rowCode)-1; r++ {
			char := string(rowCode[r])
			if r == 0 {
				if char == frontMarker {
					lPointer = 0
					rPointer = 63
					fmt.Printf("Initial F Keeping row %v through %v\n", lPointer, rPointer)
				} else {
					lPointer = 64
					rPointer = rowsCount
					fmt.Printf("Initial B Keeping row %v through %v\n", lPointer, rPointer)
				}
			} else {
				if char == frontMarker {
					rPointer = lPointer + ((rPointer - lPointer) / 2)
					// rPointer = rPointer - ((rPointer-lPointer)+1)/2
					fmt.Printf("F Keeping row %v through %v\n", lPointer, rPointer)
					// fmt.Printf("R Pointer Set To %v\n", rPointer)
				} else {
					//lPointer = (rPointer / 2) + 1
					lPointer = lPointer + (((rPointer + 1) - lPointer) / 2)
					fmt.Printf("B Keeping row %v through %v\n", lPointer, rPointer)
					// fmt.Printf("L Pointer Set To %v\n", lPointer)
				}
			}
		}
		// Final Row
		fmt.Printf("Pointers L %v, R %v\n", lPointer, rPointer)
		row = Max(lPointer, rPointer)

		llPointer := 0
		rrPointer := 0

		for c := 0; c <= len(colCode)-1; c++ {
			char := string(colCode[c])

			if c == 0 {
				if char == leftMarker {
					rrPointer = colsCount
				} else {
					llPointer = 4
					rrPointer = colsCount
				}
			} else {
				if char == leftMarker {
					rrPointer = llPointer + ((rrPointer - llPointer) / 2)
					fmt.Printf("Keeping col %v through %v\n", llPointer, rrPointer)
				} else {
					llPointer = llPointer + (((rrPointer + 1) - llPointer) / 2)
					fmt.Printf("Keeping col %v through %v\n", llPointer, rrPointer)
				}
			}
		}
		col = Max(llPointer, rrPointer)

		prod := (row * 8) + col
		best = Max(best, prod)
	}

	return best
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	data, err := ioutil.ReadFile("./input_t.txt")
	checkError(err)

	passes := strings.Split(string(data), "\n")

	fmt.Printf("Highest Seat ID %v", highestSeatID(passes))
}
