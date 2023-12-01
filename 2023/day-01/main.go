package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Printf("The answer is %d\n", Solve(os.Args[1]))
}

func Solve(inputFilename string) int {
	result := 0
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		firstAndLast := []rune{
			'a',
			'a',
		}
		number := []rune{}

		for _, v := range scanner.Text() {

			if unicode.IsDigit(v) {
				setNumber(firstAndLast, v)
				number = []rune{}
			} else {
				number = append(number, v)
				if num := sliceToNumber(number); num != '0' {
					setNumber(firstAndLast, num)
					number = number[len(number)-3:]
				}
			}
		}
		if firstAndLast[0] == 'a' {
			continue
		}
		if firstAndLast[1] == 'a' {
			firstAndLast[1] = firstAndLast[0]
		}
		realNumber, err := strconv.Atoi(string(firstAndLast))
		if err != nil {
			panic(err)
		}
		result = result + realNumber
	}

	return result
}

func setNumber(firstAndLast []rune, v rune) {
	if firstAndLast[0] == 'a' {
		firstAndLast[0] = v
	} else {
		firstAndLast[1] = v
	}
}

func sliceToNumber(number []rune) rune {
	num := '0'
	for i := 3; i <= len(number); i++ {
		switch string(number[len(number)-i:]) {
		case "one":
			return '1'
		case "two":
			return '2'
		case "three":
			return '3'
		case "four":
			return '4'
		case "five":
			return '5'
		case "six":
			return '6'
		case "seven":
			return '7'
		case "eight":
			return '8'
		case "nine":
			return '9'
		}
	}

	return num
}
