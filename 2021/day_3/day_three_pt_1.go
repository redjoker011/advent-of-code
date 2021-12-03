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

	input := strings.Split(string(data), "\n")
	var gammaRate string
	var epsilonRate string

	for i := 0; i < 12; i++ {
		var zeroes int
		var ones int

		for idx, v := range input {
			if idx == len(input)-1 {
				fmt.Println("Reached EOF \n")
				break
			}

			if string(v[i]) == "1" {
				ones += 1
			} else {
				zeroes += 1
			}
		}

		fmt.Printf("Zeroes %v; Ones %v", zeroes, ones)

		if zeroes > ones {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			epsilonRate += "0"
			gammaRate += "1"
		}
	}

	fmt.Printf("\nGamma %v; Epsilon %v \n", gammaRate, epsilonRate)

	gammaRateInt, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonRateInt, _ := strconv.ParseInt(epsilonRate, 2, 64)
	fmt.Printf("Power Consumption Rate %v", gammaRateInt*epsilonRateInt)
}
