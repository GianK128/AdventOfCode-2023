package day1

import (
	"bufio"
	"os"
	"regexp"
	"unicode"
)

// ====== PART 1 ====== //

const byteDigit0 = 48

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

		sum += int((code[first]-byteDigit0)*10 + (code[last] - byteDigit0))
	}

	return sum
}

// ====== PART 2 ====== //

var spelledNumbersValues = map[string]*regexp.Regexp{
	"one":   regexp.MustCompile("one"),
	"two":   regexp.MustCompile("two"),
	"three": regexp.MustCompile("three"),
	"four":  regexp.MustCompile("four"),
	"five":  regexp.MustCompile("five"),
	"six":   regexp.MustCompile("six"),
	"seven": regexp.MustCompile("seven"),
	"eight": regexp.MustCompile("eight"),
	"nine":  regexp.MustCompile("nine"),
}

func getSpelledNumberValue(spelledNumber string) byte {
	switch spelledNumber {
	case "one":
		return byteDigit0 + 1
	case "two":
		return byteDigit0 + 2
	case "three":
		return byteDigit0 + 3
	case "four":
		return byteDigit0 + 4
	case "five":
		return byteDigit0 + 5
	case "six":
		return byteDigit0 + 6
	case "seven":
		return byteDigit0 + 7
	case "eight":
		return byteDigit0 + 8
	case "nine":
		return byteDigit0 + 9
	}
	return 0
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

		firstChar := code[first]
		lastChar := code[last]

		// We need to check for numbers spelled out in the string
		// After checking for each number, we will store the minimum and maximum index
		minIndex := first
		maxIndex := last
		for spelledNumber, spelledNumberRegex := range spelledNumbersValues {
			matches := spelledNumberRegex.FindAllStringIndex(code, -1)
			if matches != nil {
				for _, match := range matches {
					index := match[0]
					if index < minIndex {
						minIndex = index
						firstChar = getSpelledNumberValue(spelledNumber)
					}
					if index > maxIndex {
						maxIndex = index
						lastChar = getSpelledNumberValue(spelledNumber)
					}
				}
			}
		}

		sum += int((firstChar-byteDigit0)*10 + (lastChar - byteDigit0))
	}

	return sum
}
