package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Unable to open input file:", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	increased := 0
	prevs := [3]int{}
	prevSum := 0
	sum := 0
	i := 0
	for scanner.Scan() {
		deep, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln("Unable to convert input to int:", err)
		}
		if i < 3 {
			prevs[i] = deep
			i++
			sum += deep
			continue
		}

		if prevSum == 0 {
			prevSum = sum
			continue
		}

		sum = prevs[2] + prevs[1] + deep
		if sum > prevSum {
			increased++
		}
		prevSum = sum
		prevs[0] = prevs[1]
		prevs[1] = prevs[2]
		prevs[2] = deep
	}

	sum = prevs[2] + prevs[1] + prevs[0]
	if sum > prevSum {
		increased++
	}

	fmt.Printf("The dept increased %d times\n", increased)
}
