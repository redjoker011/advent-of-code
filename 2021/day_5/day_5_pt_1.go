package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2021/day_5/input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")

	type P struct {
		x1, y1 int
		x2, y2 int
	}

	var entries = make([]P, len(input))
	vhLines := []P{}

	maxX := 0
	maxY := 0
	for i, l := range input {
		fmt.Sscanf(l, "%d,%d -> %d,%d",
			&entries[i].x1,
			&entries[i].y1,
			&entries[i].x2,
			&entries[i].y2)

		if entries[i].x1 > maxX {
			maxX = entries[i].x1
		}

		if entries[i].x2 > maxX {
			maxX = entries[i].x2
		}

		if entries[i].y1 > maxY {
			maxX = entries[i].y1
		}

		if entries[i].y2 > maxY {
			maxX = entries[i].y2
		}

		if entries[i].x1 == entries[i].x2 || entries[i].y1 == entries[i].y2 {
			vhLines = append(vhLines, entries[i])
		}
	}

	fmt.Println(vhLines)

	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}

	for _, l := range vhLines {
		if l.x1 == l.x2 {
			y1 := l.y1
			y2 := l.y2

			if y2 < y1 {
				y1 = l.y2
				y2 = l.y1
			}

			for y := y1; y <= y2; y++ {
				grid[y][l.x1]++
			}
		} else {
			x1 := l.x1
			x2 := l.x2

			if x2 < x1 {
				x1 = l.x2
				x2 = l.x1
			}

			for x := x1; x <= x2; x++ {
				grid[l.y1][x]++
			}
		}
	}

	count := 0
	for _, row := range grid {
		for _, v := range row {
			if v >= 2 {
				count++
			}
		}
	}

	fmt.Println(count)
}
