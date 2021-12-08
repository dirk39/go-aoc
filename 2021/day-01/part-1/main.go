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
	prev := 0
	for scanner.Scan() {
		deep, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln("Unable to convert input to int:", err)
		}
		if prev > 0 && deep > prev {
			increased++
		}
		prev = deep
	}

	fmt.Printf("The dept increased %d times\n", increased)
}
