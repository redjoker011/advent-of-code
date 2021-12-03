package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2021/day_3/input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var oxygenGeneratorRating string
	var scrubberRating string

	oxygenBits := input
	for i := 0; i < 12; i++ {
		var oxZeroBits []string
		var oxOneBits []string

		for _, v := range oxygenBits {
			if string(v[i]) == "1" {
				oxOneBits = append(oxOneBits, v)
			} else {
				oxZeroBits = append(oxZeroBits, v)
			}
		}

		if len(oxZeroBits) == 2 && len(oxZeroBits) > len(oxOneBits) {
			if string(oxZeroBits[0][i+1]) == "1" {
				oxygenGeneratorRating = oxZeroBits[0]
				break
			} else {
				oxygenGeneratorRating = oxZeroBits[1]
				break
			}
		} else if len(oxOneBits) == 2 && len(oxOneBits) > len(oxZeroBits) {
			if string(oxOneBits[0][i+1]) == "1" {
				oxygenGeneratorRating = oxOneBits[0]
				break
			} else {
				oxygenGeneratorRating = oxOneBits[1]
				break
			}
		} else if len(oxZeroBits) > len(oxOneBits) {
			oxygenBits = oxZeroBits
		} else {
			oxygenBits = oxOneBits
		}

		if len(oxygenBits) == 1 {
			fmt.Printf("Rating Found %v \n", oxygenBits[0])
			oxygenGeneratorRating = oxygenBits[0]
			break
		}
	}

	scrubBits := input
	for i := 0; i < 12; i++ {

		var scrubZeroBits []string
		var scrubOneBits []string

		for _, v := range scrubBits {
			if string(v[i]) == "1" {
				scrubOneBits = append(scrubOneBits, v)
			} else {
				scrubZeroBits = append(scrubZeroBits, v)
			}
		}

		if len(scrubZeroBits) == 2 && len(scrubZeroBits) < len(scrubOneBits) {
			if string(scrubZeroBits[0][i+1]) == "0" {
				scrubberRating = scrubZeroBits[0]
				break
			} else {
				scrubberRating = scrubZeroBits[1]
				break
			}
		} else if len(scrubOneBits) == 2 && len(scrubOneBits) < len(scrubZeroBits) {
			if string(scrubOneBits[0][i+1]) == "0" {
				scrubberRating = scrubOneBits[0]
				break
			} else {
				scrubberRating = scrubOneBits[1]
				break
			}
		} else if len(scrubOneBits) > len(scrubZeroBits) {
			scrubBits = scrubZeroBits
		} else if len(scrubOneBits) == len(scrubZeroBits) {
			scrubBits = scrubZeroBits
		} else {
			scrubBits = scrubOneBits
		}

		if len(scrubBits) == 1 {
			scrubberRating = scrubBits[0]
			break
		}
	}

	oxygenInt, _ := strconv.ParseInt(oxygenGeneratorRating, 2, 64)
	scrubInt, _ := strconv.ParseInt(scrubberRating, 2, 64)
	fmt.Printf("\noxygen %v %v; scrub %v %v", oxygenInt, oxygenGeneratorRating, scrubInt, scrubberRating)
	fmt.Printf("\nLife Support Rating %v", oxygenInt*scrubInt)
}
