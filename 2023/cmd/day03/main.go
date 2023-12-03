package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/eotsn/advent_of_code/2023/file"
)

func main() {
	lines, err := file.ReadLines("input/03.txt")
	if err != nil {
		panic(err)
	}

	symbols := make(map[int][]int) // store symbol indices for each line
	for i, line := range lines {
		for j, rune := range line {
			if !unicode.IsDigit(rune) && rune != '.' {
				symbols[i] = append(symbols[i], j)
			}
		}
	}

	var sum int

	for i, line := range lines {
		symbolIndices := symbols[i] // combine indices for surrounding lines
		if i != 0 {
			symbolIndices = append(symbolIndices, symbols[i-1]...)
		}
		if i != len(lines) {
			symbolIndices = append(symbolIndices, symbols[i+1]...)
		}

		var partNo string

		start := -1 // save the start index of the current part number
		for j, rune := range line {
			if unicode.IsDigit(rune) {
				if start < 0 {
					start = j
				}
				partNo += fmt.Sprintf("%c", rune)
				if j == len(line)-1 { // make sure we catch part numbers at the end of the line
					for _, symbolIndex := range symbolIndices {
						if symbolIndex >= start-1 && symbolIndex <= j {
							n, _ := strconv.Atoi(partNo)
							sum += n
						}
					}
					partNo = ""
					start = -1
				}
			} else if start >= 0 {
				for _, symbolIndex := range symbolIndices {
					if symbolIndex >= start-1 && symbolIndex <= j {
						n, _ := strconv.Atoi(partNo)
						sum += n
					}
				}
				partNo = ""
				start = -1
			}
		}
	}

	fmt.Println(sum)
}
