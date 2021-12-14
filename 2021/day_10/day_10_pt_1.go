package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2021/day_10/input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")
	normalizedInput := input[0 : len(input)-1]

	fmt.Println(normalizedInput[len(normalizedInput)-1])

	score := 0
	// Iterate Over Inputs
	for _, line := range normalizedInput {

		// Store Stack and Corrupted Character as Rune
		stack := []rune{}
		corrupted := rune(0)

	loop:
		// Iterage over array of characters
		for _, c := range line {
			fmt.Println(string(c), string(stack), "\n")
			switch c {
			case '(', '[', '{', '<':
				fmt.Println("Open", string(c), string(stack), "\n")
				stack = append(stack, c)
				fmt.Println("Append", string(c), string(stack), "\n")
			case ')', ']', '}', '>':
				fmt.Println("Close", string(c), string(stack), "\n")
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fmt.Println("Pop", string(c), string(stack), "\n")

				// Check if character is complete
				if c == ')' && pop == '(' ||
					c == ']' && pop == '[' ||
					c == '}' && pop == '{' ||
					c == '>' && pop == '<' {
					continue
				} else {
					corrupted = c
					break loop
				}
			}
		}

		switch corrupted {
		case ')':
			score += 3
		case ']':
			score += 57
		case '}':
			score += 1197
		case '>':
			score += 25137
		default:
		}
	}

	fmt.Println(score)
}
