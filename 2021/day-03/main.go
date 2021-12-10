package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	zero int
	one  int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Cannot open input.txt. Err ", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	// Part 1
	partOne(scanner)
	file.Close()
	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatalln("Cannot open input.txt. Err ", err)
	}
	scanner = bufio.NewScanner(file)
}

func partOne(scanner *bufio.Scanner) map[int]*point {
	fmt.Println("----- Part One -----")
	pointStats := map[int]*point{}
	for scanner.Scan() {
		line := scanner.Bytes()
		for i, v := range line {
			if _, ok := pointStats[i]; !ok {
				pointStats[i] = &point{}
			}
			switch v {
			case '0':
				pointStats[i].zero++
			case '1':
				pointStats[i].one++
			}
		}
	}
	gamma := 0
	y := 0
	pos := len(pointStats) - 1
	for i, point := range pointStats {
		if point.one > point.zero {
			fmt.Printf("most common 1 << %d - %d = %d\n", pos, i, 1<<(pos-i))
			gamma += 1 << (pos - i)
		} else {
			fmt.Printf("less common 1 << %d - %d = %d\n", pos, i, 1<<(pos-i))
			y += 1 << (pos - i)
		}
	}
	fmt.Printf("Gamma: %d\n", gamma)
	fmt.Printf("Y: %d\n", y)
	fmt.Printf("Checksum: %d\n", gamma*y)
	fmt.Println("----- End Part One -----")

	return pointStats
}

func partTwo(scanner *bufio.Scanner) {
	fmt.Println("----- Part Two -----")

	fmt.Println("----- End Part Two -----")
}
