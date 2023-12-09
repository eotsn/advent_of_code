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

	var startNodes []string

	nodes := make(map[string][]string)
	for _, line := range lines[2:] { // skip first two lines
		var (
			node  = line[:3]
			left  = line[7:10]
			right = line[12:15]
		)
		if node[2] == 'A' {
			startNodes = append(startNodes, node)
		}
		nodes[node] = []string{left, right}
	}

	var steps []int
	for _, node := range startNodes {
		currentNode := node

		var i, j int
		for {
			if currentNode[2] == 'Z' {
				steps = append(steps, j)
				break
			}
			switch instructions[i] {
			case 'L':
				currentNode = nodes[currentNode][0]
			case 'R':
				currentNode = nodes[currentNode][1]
			}
			j++
			i = (i + 1) % len(instructions)
		}
	}

	fmt.Println(lcm(steps[0], steps[1], steps[2:]...))
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)
	for _, i := range integers {
		result = lcm(result, i)
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
