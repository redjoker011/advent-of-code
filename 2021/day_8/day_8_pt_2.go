package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2021/day_8/input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")

	sum := 0
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

		digits := deduce(strings.Split(v2[0], " "))
		output := digits[normalize(v3[0])]*1000 +
			digits[normalize(v3[1])]*100 +
			digits[normalize(v3[2])]*10 +
			digits[normalize(v3[3])]

		sum += output
	}
	fmt.Println(count, sum)
}

func deduce(digits []string) map[string]int {
	var len2, len3, len4, len7 string
	var len5 []string
	var len6 []string

	for _, v := range digits {
		if len(v) == 2 {
			len2 = normalize(v)
		} else if len(v) == 3 {
			len3 = normalize(v)
		} else if len(v) == 4 {
			len4 = normalize(v)
		} else if len(v) == 7 {
			len7 = normalize(v)
		} else if len(v) == 5 {
			len5 = append(len5, normalize(v))
		} else {
			len6 = append(len6, normalize(v))
		}
	}

	var out = map[string]int{
		len2: 1,
		len4: 4,
		len3: 7,
		len7: 8,
	}

	var digit6 string
	// find 6
	for i, d := range len6 {
		if strings.Index(d, string(len2[0])) == -1 {
			out[d] = 6
			digit6 = d
			len6 = append(len6[:i], len6[i+1:]...)
			break
		}
		if strings.Index(d, string(len2[1])) == -1 {
			out[d] = 6
			digit6 = d
			len6 = append(len6[:i], len6[i+1:]...)
			break
		}
	}

	// find 0
	for i, d := range len6 {
		if strings.Index(d, string(len4[0])) == -1 {
			out[d] = 0
			len6 = append(len6[:i], len6[i+1:]...)
			break
		}
		if strings.Index(d, string(len4[1])) == -1 {
			out[d] = 0
			len6 = append(len6[:i], len6[i+1:]...)
			break
		}
		if strings.Index(d, string(len4[2])) == -1 {
			out[d] = 0
			len6 = append(len6[:i], len6[i+1:]...)
			break
		}
		if strings.Index(d, string(len4[3])) == -1 {
			out[d] = 0
			len6 = append(len6[:i], len6[i+1:]...)
			break
		}
	}

	// remaining is 9
	out[len6[0]] = 9

	// find 3: contains all of 1
	for i, d := range len5 {
		if strings.Index(d, string(len2[0])) != -1 && strings.Index(d, string(len2[1])) != -1 {
			out[d] = 3
			len5 = append(len5[:i], len5[i+1:]...)
			break
		}
	}

	// find 5: is all in 6
	for i, d := range len5 {
		if strings.Index(digit6, string(d[0])) != -1 &&
			strings.Index(digit6, string(d[1])) != -1 &&
			strings.Index(digit6, string(d[2])) != -1 &&
			strings.Index(digit6, string(d[3])) != -1 &&
			strings.Index(digit6, string(d[4])) != -1 {
			out[d] = 5
			len5 = append(len5[:i], len5[i+1:]...)
			break
		}
	}

	// remaining is 2
	out[len5[0]] = 2

	return out
}

// Sort String in Alphabetical Order
func normalize(digit string) string {
	arr := strings.Split(digit, "")
	sort.Strings(arr)
	return strings.Join(arr, "")
}
