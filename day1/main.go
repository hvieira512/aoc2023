package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var NUMBERS = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func getPart1(data []byte) int {
	result := 0
	numberLineList := []string{}

	for _, char := range string(data) {
		if unicode.IsDigit(char) {
			numberLineList = append(numberLineList, string(char))
		}

		if unicode.IsSpace(char) {
			lastPos := len(numberLineList) - 1
			strAux := string(numberLineList[0]) + string(numberLineList[lastPos])

			lineNumber, err := strconv.Atoi(strAux)
			if err != nil {
				log.Fatal(err)
			}
			result = result + lineNumber

			// Reset for next line
			numberLineList = []string{}
		}
	}

	return result
}

func getPart2(data []byte) int {
	result := 0
	numberLineList := []string{}

	for _, char := range string(data) {
		if unicode.IsDigit(char) {
			numberLineList = append(numberLineList, string(char))
		}

		if unicode.IsSpace(char) {
			lastPos := len(numberLineList) - 1
			strAux := string(numberLineList[0]) + string(numberLineList[lastPos])

			lineNumber, err := strconv.Atoi(strAux)
			if err != nil {
				log.Fatal(err)
			}
			result = result + lineNumber

			// Reset for next line
			numberLineList = []string{}
		}
	}

	return result
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	resultPart1 := getPart1(data)
	fmt.Printf("Part 1: %d\n", resultPart1)

	resultPart2 := getPart2(data)
	fmt.Printf("Part 2: %d\n", resultPart2)
}
