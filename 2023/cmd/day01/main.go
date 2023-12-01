package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/eotsn/advent_of_code/2023/file"
)

// Because golang's regex implementation (RE2) doesn't support positive
// lookahead we have to do this the hard way.
var regex = regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine")

var digits = map[string]string{
	"0":     "0",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	lines, err := file.ReadLines("input/01.txt")
	if err != nil {
		panic(err)
	}

	var sum int
	for _, line := range lines {
		var num, s string

		s = ""
		for i := 0; i < len(line); i++ {
			s = fmt.Sprintf("%s%c", s, line[i])
			if m := regex.FindString(s); m != "" {
				num += digits[m]
				break
			}
		}

		s = ""
		for i := len(line) - 1; i >= 0; i-- {
			s = fmt.Sprintf("%c%s", line[i], s)
			if m := regex.FindString(s); m != "" {
				num += digits[m]
				break
			}
		}

		if n, err := strconv.Atoi(num); err == nil {
			sum += n
		}
	}

	fmt.Println(sum)
}
