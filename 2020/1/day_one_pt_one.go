package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input_one.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")
	in := 0
	inVal, _ := strconv.Atoi(input[in])
	ans := 0

	for ans == 0 {
		for idx, v := range input {
			intV, _ := strconv.Atoi(v)

			fmt.Printf("Sum %v of %v and %v \n", inVal+intV, inVal, intV)
			if idx == len(input)-1 {
				fmt.Println("Reached EOF")
				inVal, _ = strconv.Atoi(input[in+1])
				in++
			} else if inVal+intV == 2020 {
				fmt.Printf("Answer %v", inVal*intV)
				ans = inVal * intV
				break
			}
		}
	}
}
