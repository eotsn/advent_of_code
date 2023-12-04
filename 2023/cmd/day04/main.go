package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/eotsn/advent_of_code/2023/file"
)

func main() {
	lines, err := file.ReadLines("input/04.txt")
	if err != nil {
		panic(err)
	}

	var sum int

	var s []string
	for _, line := range lines {
		w := make(map[int]bool) // assemble winning numbers

		s = strings.Split(line, ":")
		s = strings.Split(s[1], "|")

		for _, n := range parseInts(s[0]) {
			w[n] = true
		}

		count := -1.0
		for _, n := range parseInts(s[1]) {
			if w[n] {
				count++
			}
		}

		if count >= 0 {
			sum += int(math.Pow(2, count))
		}
	}

	fmt.Println(sum)
}

func parseInts(s string) []int {
	r := regexp.MustCompile(`(\d+)`)
	matches := r.FindAllString(s, -1)

	nums := make([]int, len(matches))
	for i, match := range matches {
		nums[i], _ = strconv.Atoi(match)
	}
	return nums
}
