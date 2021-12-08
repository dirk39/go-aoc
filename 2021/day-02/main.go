package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Could not read file:", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	partOne(scanner)
	partTwo(scanner)
}

func partOne(scanner *bufio.Scanner) {
	fmt.Println("----- Part One -----")
	var forward, dept int
	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")
		distance, err := strconv.Atoi(commands[1])
		if err != nil {
			log.Fatalln("Could not convert distance to int:", err)
		}
		switch commands[0] {
		case "forward":
			forward += distance
		case "down":
			dept += distance
		case "up":
			dept -= distance
		}
	}
	fmt.Printf("The forward distance is %d\n", forward)
	fmt.Printf("The dept is %d\n", dept)
	fmt.Printf("The total distance is %d\n", forward*dept)
	fmt.Println("----- End Part One -----")
}

func partTwo(scanner *bufio.Scanner) {
	fmt.Println("----- Part Two -----")
	var forward, dept, aim int
	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")
		distance, err := strconv.Atoi(commands[1])
		if err != nil {
			log.Fatalln("Could not convert distance to int:", err)
		}
		switch commands[0] {
		case "forward":
			forward += distance
			dept += distance * aim
		case "down":
			aim += distance
		case "up":
			aim -= distance
		}
	}
	fmt.Printf("The forward distance is %d\n", forward)
	fmt.Printf("The dept is %d\n", dept)
	fmt.Printf("The total distance is %d\n", forward*dept)
	fmt.Println("----- End Part Two -----")
}
