package main

import (
	"fmt"

	"github.com/eotsn/advent_of_code/2023/file"
)

func main() {
	lines, err := file.ReadLines("input/08.txt")
	if err != nil {
		panic(err)
	}

	instructions := lines[0]

	nodes := make(map[string][]string)
	for _, line := range lines[2:] { // skip first two lines
		var (
			node  = line[:3]
			left  = line[7:10]
			right = line[12:15]
		)
		nodes[node] = []string{left, right}
	}

	var (
		steps       = 0
		currentNode = "AAA"
	)

	var i = 0
	for {
		if currentNode == "ZZZ" {
			fmt.Println(fmt.Sprintf("found ZZZ in %d steps", steps))
			break
		}
		switch instructions[i] {
		case 'L':
			currentNode = nodes[currentNode][0]
		case 'R':
			currentNode = nodes[currentNode][1]
		}
		steps++
		i = (i + 1) % len(instructions)
	}
}
