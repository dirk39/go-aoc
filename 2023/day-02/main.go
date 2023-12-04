package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part1, part2 := Solve("input.txt")
	fmt.Printf("The answer is %d %d\n", part1, part2)
}

func Solve(inputFilename string) (int, int) {
	result1 := 0
	result2 := 0
	gameCount := 1
	colorx, _ := regexp.Compile(`(\d+) (\w+)`)
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		min := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		line := scanner.Text()
		colorMatch := colorx.FindAllStringSubmatch(line, -1)
		for _, colors := range colorMatch {
			num, _ := strconv.Atoi(colors[1])
			if num > min[colors[2]] {
				min[colors[2]] = num
			}
		}

		if min["red"] <= 12 && min["green"] <= 13 && min["blue"] <= 14 {
			result1 += gameCount
		}

		result2 += min["red"] * min["green"] * min["blue"]
		gameCount = gameCount + 1
	}

	return result1, result2
}
