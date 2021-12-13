package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2021/day_9/input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")

	var grid [][]int = make([][]int, len(input)-1)
	for i, row := range input {

		if i == len(input)-1 {
			break
		}

		grid[i] = make([]int, len(row))

		for j, v := range row {
			var err error
			grid[i][j], err = strconv.Atoi(string(v))

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	maxX := len(grid[0])
	maxY := len(grid)

	risk := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Println("grid", grid[y], "\n")
			curr := grid[y][x]

			// Check Up
			if y > 0 && grid[y-1][x] <= curr {
				fmt.Printf("Up Y Index %v; X Index %v; Value %v; Next %v \n", y, x, curr, grid[y-1][x])
				continue
			}

			// Check Down
			if y < maxY-1 && grid[y+1][x] <= curr {
				fmt.Printf("Down Y Index %v; X Index %v; Value %v; Next %v \n", y, x, curr, grid[y+1][x])
				continue
			}

			// Check Left
			if x > 0 && grid[y][x-1] <= curr {
				fmt.Printf("L Y Index %v; X Index %v; Value %v; Next %v \n", y, x, curr, grid[y][x-1])
				continue
			}

			// Check Right
			if x < maxX-1 && grid[y][x+1] <= curr {
				fmt.Printf("R Y Index %v; X Index %v; Value %v; Next %v \n", y, x, curr, grid[y][x+1])
				continue
			}

			risk += curr + 1
		}
	}
	fmt.Println("\nRisk", risk)
}
