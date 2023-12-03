package main

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"

	"github.com/eotsn/advent_of_code/2023/file"
)

type Gear struct {
	x, y int
}

type PartNo struct {
	value int
	x0    int
	x1    int
}

func main() {
	lines, err := file.ReadLines("input/03.txt")
	if err != nil {
		panic(err)
	}

	// add some padding so we can skip edge cases for the 1st and last lines
	var padding string
	for i := 0; i < len(lines[0]); i++ {
		padding += "."
	}
	lines = append([]string{padding}, lines...)
	lines = append(lines, padding)

	var gears []Gear
	parts := make([][]PartNo, len(lines)) // store part numbers for each line
	for i, line := range lines {
		for j, rune := range line {
			if !unicode.IsDigit(rune) && rune == '*' {
				gears = append(gears, Gear{y: i, x: j})
			}
		}

		r := regexp.MustCompile(`(\d+)`)
		matches := r.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			partNo, _ := strconv.Atoi(line[match[0]:match[1]])
			parts[i] = append(parts[i], PartNo{
				value: partNo,
				x0:    match[0],
				x1:    match[1],
			})
		}
	}

	var sum int

	for _, gear := range gears {
		var p []int

		// collect the matching part numbers for the previous, current and next line
		for _, y := range []int{gear.y - 1, gear.y + 1, gear.y} {
			for _, part := range parts[y] {
				if part.x0-1 <= gear.x && part.x1 >= gear.x {
					p = append(p, part.value)
				}
			}
		}
		if len(p) == 2 {
			sum += p[0] * p[1]
		}
	}

	fmt.Println(sum)
}
