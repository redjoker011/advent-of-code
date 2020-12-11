package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := os.Open("./input_two.txt")
	checkError(err)

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		eachNum, err := strconv.Atoi(scanner.Text())
		checkError(err)
		numbers = append(numbers, eachNum)
	}

	sum := 2020
	sort.Ints(numbers)
	fmt.Printf("Sorted Numbers %v \n", numbers)

	ans, err := find(numbers, sum)
	checkError(err)
	fmt.Println("Answer: ", product(ans))
}

func product(array []int) int {
	result := 1
	for _, eachNum := range array {
		result *= eachNum
	}
	return result
}

func find(numbers []int, sum int) ([]int, error) {
	endNum := len(numbers) - 1
	for i := 0; i < len(numbers)-1; i++ {
		innerStartNum := 0
		innerEndNum := endNum - 1
		for j := 0; j < (endNum - 1); j++ {
			tempSum := numbers[innerStartNum] + numbers[innerEndNum] + numbers[endNum]
			fmt.Printf("Temporary Sum %v in index %v-%v-%v \n", tempSum, innerStartNum, innerEndNum, endNum)
			if tempSum > sum {
				innerEndNum -= 1
			}
			if tempSum < sum {
				innerStartNum += 1
			}
			if tempSum == sum {
				threeNumbers := make([]int, 3, 3)
				threeNumbers[0] = numbers[innerStartNum]
				threeNumbers[1] = numbers[innerEndNum]
				threeNumbers[2] = numbers[endNum]
				return threeNumbers, nil
			}
		}
		endNum -= 1
	}
	return nil, fmt.Errorf("Unable to find match")
}
