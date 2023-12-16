package day1

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

// ====== PART 1 ====== //

func Part1(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code := scanner.Text()

		// Index of first and last character
		first := 0
		last := len(code) - 1

		// Loop through the string until the character is a number
		for !unicode.IsDigit(rune(code[first])) {
			first++
		}

		// Loop through the string until the character is a number
		for !unicode.IsDigit(rune(code[last])) {
			last--
		}

		calibrationCode := code[first:first+1] + code[last:last+1]
		calibrationCodeInt, _ := strconv.Atoi(calibrationCode)
		sum += calibrationCodeInt
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

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code := scanner.Text()

		first := 0
		last := len(code) - 1

		// Loop through the string until the character is a number
		for !unicode.IsDigit(rune(code[first])) {
			first++
		}

		// Loop through the string until the character is a number
		for !unicode.IsDigit(rune(code[last])) {
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
		sum += calibrationCodeInt
	}

	return sum
}
