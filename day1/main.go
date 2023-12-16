package day1

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

// ====== UTILS ====== //

func isNumber(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}

// ====== PART 1 ====== //

func Part1(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read file line by line and store in codes
	codes := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}

	calibrationCodes := []int{}
	for _, code := range codes {
		// Index of first and last character
		first := 0
		last := len(code) - 1

		// Loop through the string until the character is a number
		for !isNumber(code[first : first+1]) {
			first++
		}

		// Loop through the string until the character is a number
		for !isNumber(code[last : last+1]) {
			last--
		}

		calibrationCode := code[first:first+1] + code[last:last+1]
		calibrationCodeInt, _ := strconv.Atoi(calibrationCode)
		calibrationCodes = append(calibrationCodes, calibrationCodeInt)
	}

	// Sum all calibration codes
	sum := 0
	for _, calibrationCode := range calibrationCodes {
		sum += calibrationCode
	}

	return sum
}

// ====== PART 2 ====== //

var possibleSpelledNumbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var spelledNumbersValues = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func Part2(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	codes := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}

	calibrationCodes := []int{}
	for _, code := range codes {
		first := 0
		last := len(code) - 1

		// Loop through the string until the character is a number
		for !isNumber(code[first : first+1]) {
			first++
		}

		// Loop through the string until the character is a number
		for !isNumber(code[last : last+1]) {
			last--
		}

		firstChar := code[first : first+1]
		lastChar := code[last : last+1]

		// We need to check for numbers spelled out in the string
		// After checking for each number, we will store the minimum and maximum index
		minIndex := first
		maxIndex := last
		for _, spelledNumber := range possibleSpelledNumbers {
			r := regexp.MustCompile(spelledNumber)
			matches := r.FindAllStringIndex(code, -1)
			if matches != nil && len(matches) > 0 {
				for _, match := range matches {
					index := match[0]
					if index < minIndex {
						minIndex = index
						firstChar = spelledNumbersValues[spelledNumber]
					}
					if index > maxIndex {
						maxIndex = index
						lastChar = spelledNumbersValues[spelledNumber]
					}
				}
			}
		}

		calibrationCode := firstChar + lastChar
		calibrationCodeInt, _ := strconv.Atoi(calibrationCode)
		calibrationCodes = append(calibrationCodes, calibrationCodeInt)
	}

	// Sum all calibration codes
	sum := 0
	for _, calibrationCode := range calibrationCodes {
		sum += calibrationCode
	}

	return sum
}
