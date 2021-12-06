package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2021/day_4/input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")
	randNumbers := strings.Split(input[0], ",")

	var numbers []int

	for _, n := range randNumbers {
		i, err := strconv.Atoi(n)

		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, i)
	}

	var boards [][]int
	for i := 2; i < len(input); i += 6 {
		var board []int

		for _, s := range strings.Split(strings.Join(input[i:i+5], " "), " ") {
			if s == "" {
				continue
			}

			i, err := strconv.Atoi(s)

			if err != nil {
				log.Fatal(err)
			}

			board = append(board, i)
		}

		if len(board) != 25 {
			log.Fatal(board)
		}

		boards = append(boards, board)
	}

	boardWin := make([]bool, len(boards))

	for _, n := range numbers {
		for b := range boards {
			if boardWin[b] {
				continue
			}

			for i, v := range boards[b] {
				if n == v {
					boards[b][i] = 0
					break
				}
			}
			// Check for winning board

			if checkWin(boards[b]) {
				sum := 0
				for _, j := range boards[b] {
					sum += j
				}

				fmt.Println(n, boards[b], sum*n)
				boardWin[b] = true
			}
		}
	}
}

func checkWin(b []int) bool {
	win := true
	for i := 0; i < 5; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 5; i < 10; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 10; i < 15; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 15; i < 20; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 20; i < 25; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	// Check Rows
	win = true
	for i := 0; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 1; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 2; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 3; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 4; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	return false
}
