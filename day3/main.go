package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type PartNumber struct {
	Number     int
	StartIndex int
	EndIndex   int
}

func Part1(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Key: Number of row the part is in
	// Value: List of part numbers in that row
	schematicNumbers := map[int][]PartNumber{}

	// Key: Number of row
	// Value: Range of columns that are to be summed
	sumRanges := map[int][][]int{}

	// Regex to find numbers between dots
	numberRegex := regexp.MustCompile(`\.?(\d+)\.?`)

	// Regex to find symbols between dots
	symbolRegex := regexp.MustCompile(`\.?\d?[*$#@/%&\=\\\+\-]\d?\.?`)

	totalSum := 0
	currentRow := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := numberRegex.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			// Get the number
			number := 0
			fmt.Sscanf(line[match[2]:match[3]], "%d", &number)

			// Store the number
			schematicNumbers[currentRow] = append(schematicNumbers[currentRow], PartNumber{
				Number:     number,
				StartIndex: match[2],
				EndIndex:   match[3] - 1,
			})
		}

		symbolMatches := symbolRegex.FindAllStringIndex(line, -1)
		for _, match := range symbolMatches {
			// Store the range of columns to be summed
			sumRanges[currentRow-1] = append(sumRanges[currentRow-1], []int{match[0], match[1] - 1})
			sumRanges[currentRow] = append(sumRanges[currentRow], []int{match[0], match[1] - 1})
			sumRanges[currentRow+1] = append(sumRanges[currentRow+1], []int{match[0], match[1] - 1})
		}

		currentRow++
	}

	// Numbers that must be summed are the ones that have the start or end in range
	for i := 0; i < currentRow; i++ {
		for _, partNumber := range schematicNumbers[i] {
			for _, sumRange := range sumRanges[i] {
				// Either the start or the end of the number is in range, or the number is in range
				if (partNumber.StartIndex <= sumRange[0] && partNumber.EndIndex >= sumRange[0]) || (partNumber.StartIndex <= sumRange[1] && partNumber.EndIndex >= sumRange[1]) || (partNumber.StartIndex >= sumRange[0] && partNumber.EndIndex <= sumRange[1]) {
					totalSum += partNumber.Number
					break
				}
			}
		}
	}

	// Es entre 528784 y 536110
	return totalSum
}
