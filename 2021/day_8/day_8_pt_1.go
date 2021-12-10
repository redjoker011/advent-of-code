package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2021/day_8/input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")

	count := 0
	for idx, v := range input {
		if idx >= len(input)-1 {
			break
		}

		v2 := strings.Split(v, " | ")
		v3 := strings.Split(v2[1], " ")

		for _, d := range v3 {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				count++
			}
		}
	}

	fmt.Println(count)
}
