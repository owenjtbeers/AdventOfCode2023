package sol1

import (
	"regexp"
	"strconv"

	mylib "aoc/lib"
)

func Sol1() int {
	// Sometime we should be able to
	scanner := mylib.ReadLines("./solutions/1/input.txt")

	// Read and parse each line
	sum := 0
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line here
		sum += calculateStringValue(line, lineNumber)
		lineNumber++
	}
	return sum
}

var replacements = map[string]string{
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

func calculateStringValue(s string, lineNumber int) int {
	// Calculate the value of the string
	digitRegex := regexp.MustCompile(`\d`)
	matches := digitRegex.FindAllStringIndex(s, -1)

	maxIndex, minIndex := 0, 0
	var first, second string
	if len(matches) > 0 {
		maxIndex = matches[len(matches)-1][0]
		minIndex = matches[0][0]
		first = string(s[minIndex])
		second = string(s[maxIndex])
	}

	for old, new := range replacements {
		matches = regexp.MustCompile(old).FindAllStringIndex(s, -1)
		if len(matches) > 0 {
			if matches[0][0] <= minIndex {
				minIndex = matches[0][0]
				first = new
			}
			if matches[len(matches)-1][0] >= maxIndex {
				maxIndex = matches[len(matches)-1][0]
				second = new
			}
		}
	}

	combination := (first + second)
	convertedToInt, _ := strconv.Atoi(combination)
	return convertedToInt
}
