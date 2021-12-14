package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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

	var grid2 [][]rune = make([][]rune, len(input)-1)
	for i, row := range input {
		if i == len(input)-1 {
			break
		}
		grid2[i] = make([]rune, len(row))
	}

	name := 'A'
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid2[y][x] != 0 {
				continue
			}

			if grid[y][x] == 9 {
				grid2[y][x] = '9'
				continue
			}

			type Point struct {
				x, y int
			}

			queue := []Point{Point{x, y}}

			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]

				grid2[p.y][p.x] = name

				if p.y > 0 && grid[p.y-1][p.x] != 9 && grid2[p.y-1][p.x] == 0 {
					queue = append(queue, Point{p.x, p.y - 1})
				}

				if p.y < maxY-1 && grid[p.y+1][p.x] != 9 && grid2[p.y+1][p.x] == 0 {
					queue = append(queue, Point{p.x, p.y + 1})
				}

				if p.x > 0 && grid[p.y][p.x-1] != 9 && grid2[p.y][p.x-1] == 0 {
					queue = append(queue, Point{p.x - 1, p.y})
				}

				if p.x < maxX-1 && grid[p.y][p.x+1] != 9 && grid2[p.y][p.x+1] == 0 {
					queue = append(queue, Point{p.x + 1, p.y})
				}
			}

			name++
		}
	}

	risk := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
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

	counts := map[rune]int{}

	for y := 0; y < len(grid2); y++ {
		for x := 0; x < len(grid2[y]); x++ {
			if grid2[y][x] == '9' {
				continue
			}

			cur := grid2[y][x]

			counts[cur]++
		}
	}

	var out []int

	for _, v := range counts {
		out = append(out, v)
	}
	sort.Ints(out)
	fmt.Println("\nRisk", risk, out)
}
