package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/eotsn/advent_of_code/2023/file"
)

func main() {
	lines, err := file.ReadLines("input/02.txt")
	if err != nil {
		panic(err)
	}

	var powerSum int
	for _, line := range lines {
		var sc scanner.Scanner
		sc.Init(strings.NewReader(line))

		cubes := make(map[string]int) // defaults all values to 0

		var prevToken string
		for token := sc.Scan(); token != scanner.EOF; token = sc.Scan() {
			text := sc.TokenText()

			switch text {
			case "red", "green", "blue":
				count, _ := strconv.Atoi(prevToken)
				if count > cubes[text] {
					cubes[text] = count
				}
			}
			prevToken = text
		}

		sum := 1
		for _, count := range cubes {
			sum *= count // this will break if we have a game where any color is never picked
		}
		powerSum += sum
	}

	fmt.Println(powerSum)
}
