package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const TOTAL_RED_BOXES = 12
const TOTAL_GREEN_BOXES = 13
const TOTAL_BLUE_BOXES = 14

// ===== PART 1 ===== //

func Part1(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sumOfIDs := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameStrings := strings.Split(scanner.Text(), ":")
		gameID, _ := strconv.Atoi(gameStrings[0][5:])
		gameRolls := strings.Split(gameStrings[1], ";")
		gameWasPossible := true
		for _, roll := range gameRolls {
			roll = strings.ReplaceAll(roll, " ", "")
			boxes := strings.Split(roll, ",")
			for _, box := range boxes {
				switch {
				case strings.Contains(box, "red"):
					box = strings.Replace(box, "red", "", 1)
					boxQuantity, _ := strconv.Atoi(box)
					if boxQuantity > TOTAL_RED_BOXES {
						gameWasPossible = false
						break
					}
				case strings.Contains(box, "green"):
					box = strings.Replace(box, "green", "", 1)
					boxQuantity, _ := strconv.Atoi(box)
					if boxQuantity > TOTAL_GREEN_BOXES {
						gameWasPossible = false
						break
					}
				case strings.Contains(box, "blue"):
					box = strings.Replace(box, "blue", "", 1)
					boxQuantity, _ := strconv.Atoi(box)
					if boxQuantity > TOTAL_BLUE_BOXES {
						gameWasPossible = false
						break
					}
				}

				if !gameWasPossible {
					break
				}
			}
		}

		if gameWasPossible {
			sumOfIDs += gameID
		}
	}

	return sumOfIDs
}

// ===== PART 2 ===== //

func Part2(inputFilePath string) int {
	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sumOfPowers := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		requiredRedBoxes := 0
		requiredGreenBoxes := 0
		requiredBlueBoxes := 0

		gameStrings := strings.Split(scanner.Text(), ":")
		gameRolls := strings.Split(gameStrings[1], ";")
		for _, roll := range gameRolls {
			roll = strings.ReplaceAll(roll, " ", "")
			boxes := strings.Split(roll, ",")
			for _, box := range boxes {
				switch {
				case strings.Contains(box, "red"):
					box = strings.Replace(box, "red", "", 1)
					boxQuantity, _ := strconv.Atoi(box)
					if boxQuantity > requiredRedBoxes {
						requiredRedBoxes = boxQuantity
					}
				case strings.Contains(box, "green"):
					box = strings.Replace(box, "green", "", 1)
					boxQuantity, _ := strconv.Atoi(box)
					if boxQuantity > requiredGreenBoxes {
						requiredGreenBoxes = boxQuantity
					}
				case strings.Contains(box, "blue"):
					box = strings.Replace(box, "blue", "", 1)
					boxQuantity, _ := strconv.Atoi(box)
					if boxQuantity > requiredBlueBoxes {
						requiredBlueBoxes = boxQuantity
					}
				}
			}
		}

		sumOfPowers += requiredRedBoxes * requiredGreenBoxes * requiredBlueBoxes
	}

	return sumOfPowers
}
