package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eotsn/advent_of_code/2023/file"
)

func main() {
	lines, err := file.ReadLines("input/06.txt")
	if err != nil {
		panic(err)
	}

	t := parseNumber(lines[0])
	d := parseNumber(lines[1])

	var speed int
	for {
		speed += 1

		if distance(speed, t) > d {
			break
		}

		if speed == t {
			break
		}
	}
	fmt.Println((t - speed) - speed + 1)
}

func distance(s, t int) int {
	return s * (t - s)
}

func parseNumber(line string) int {
	s := strings.Replace(line, " ", "", -1)
	t := strings.Split(s, ":")[1]
	n, _ := strconv.Atoi(t)
	return n
}
