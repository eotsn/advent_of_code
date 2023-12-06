package main

import (
	"fmt"

	"github.com/eotsn/advent_of_code/2023/file"
	"github.com/eotsn/advent_of_code/2023/parser"
)

func main() {
	lines, err := file.ReadLines("input/06.txt")
	if err != nil {
		panic(err)
	}

	times := parser.ParseInts(lines[0])
	distances := parser.ParseInts(lines[1])

	sum := 1
	for i := 0; i < len(times); i++ {
		var speed int
		for {
			speed += 1

			if distance(speed, times[i]) > distances[i] {
				break
			}

			if speed == times[i] {
				break
			}
		}
		sum *= (times[i] - speed) - speed + 1
	}
	fmt.Println(sum)
}

func distance(s, t int) int {
	return -s*s + s*t
}
