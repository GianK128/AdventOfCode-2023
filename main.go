package main

import (
	"fmt"

	"github.com/GianK128/AdventOfCode-2023/day1"
	"github.com/GianK128/AdventOfCode-2023/day2"
	"github.com/GianK128/AdventOfCode-2023/day3"
)

func main() {
	fmt.Println("[Advent of Code 2023]")
	fmt.Println("[Day 1 - Part 1] Sum of calibration codes:", day1.Part1("day1/input.txt"))
	fmt.Println("[Day 1 - Part 2] Sum of calibration codes:", day1.Part2("day1/input.txt"))
	fmt.Println("[Day 2 - Part 1] Sum of IDs of possible games:", day2.Part1("day2/input.txt"))
	fmt.Println("[Day 2 - Part 2] Sum of the powers of set of cubes:", day2.Part2("day2/input.txt"))
	fmt.Println("[Day 3 - Part 1] Sum of adjacent part numbers:", day3.Part1("day3/input.txt"))
	fmt.Println("[Day 3 - Part 2] Sum of gear ratios:", day3.Part2("day3/input.txt"))
}
