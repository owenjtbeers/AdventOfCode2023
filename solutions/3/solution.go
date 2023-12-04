package sol3

import (
	mylib "aoc/lib"
	"fmt"
	"regexp"
	"strconv"
)

func pt1Sol3() int {

	// Read in lines
	scanner := mylib.ReadLines("./solutions/3/input.txt")
	sum := 0
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	prevLine := ""
	for i := 0; i < len(lines); i++ {
		currentLine := lines[i]

		var nextLine string
		if (i + 1) < len(lines) {
			nextLine = lines[i+1]
		} else {
			nextLine = ""
		}
		var partNumbers []int
		currentMatchIndexes := findAllDigitStrings(currentLine)
		for _, matchIndexes := range currentMatchIndexes {
			hasAdjacentSymbol := checkNeighborsForAdjacentSymbol(matchIndexes, currentLine, nextLine, prevLine)
			if hasAdjacentSymbol {
				partNumber, _ := strconv.Atoi(currentLine[matchIndexes[0]:matchIndexes[1]])
				partNumbers = append(partNumbers, partNumber)
				sum += partNumber
			}
		}
		fmt.Println(i, partNumbers)
		prevLine = currentLine
	}

	return sum
}

func Sol3() int {
	// Read in lines
	scanner := mylib.ReadLines("./solutions/3/input.txt")
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sum := 0
	prevLine := ""
	for i := 0; i < len(lines); i++ {
		currentLine := lines[i]
		for j := 0; j < len(lines[i]); j++ {
			currentChar := string(currentLine[j])
			if currentChar == "*" {
				leftBound := max(j-1, 0)
				rightBound := min(j+1, len(currentLine)-1)
				adjacentDigitMatches := findAllDigitStrings(currentLine[leftBound:rightBound])
				prevLineAdjacentDigitMatches := findAllDigitStrings(prevLine[leftBound:rightBound])
				nextLineAdjacentDigitMatches := findAllDigitStrings(lines[i+1][leftBound:rightBound])
				if len(adjacentDigitMatches)+len(prevLineAdjacentDigitMatches)+len(nextLineAdjacentDigitMatches) == 2 {
					// gearRatio := 1
					sum += 1
				}
			}
		}
	}
	return sum

}

func findAllDigitStrings(line string) [][]int {
	if line != "" {
		regex := regexp.MustCompile(`\d+`)
		stringIndexes := regex.FindAllStringIndex(line, -1)
		return stringIndexes
	}
	return [][]int{}
}

func findAllGears(line string) [][]int {
	regex := regexp.MustCompile(`[*]`)
	stringIndexes := regex.FindAllStringIndex(line, -1)
	return stringIndexes
}

func checkNeighborsForAdjacentSymbol(matchIndexes []int, currentLine string, nextLine string, prevLine string) bool {
	leftBound := max(matchIndexes[0]-1, 0)
	rightBound := min(matchIndexes[1], len(currentLine)-1)
	hasAdjacentSymbol := false
	// Check current line for adjacent symbol
	if leftBound != matchIndexes[0] {
		if checkIfSymbol(currentLine[leftBound]) {
			hasAdjacentSymbol = true
		}
	}
	if checkIfSymbol(currentLine[rightBound]) {
		hasAdjacentSymbol = true
	}
	linesToCheck := []string{prevLine, nextLine}
	// Check next line for adjacent symbol
	for _, line := range linesToCheck {
		if line != "" {
			for i := leftBound; i <= rightBound; i++ {
				if checkIfSymbol(line[i]) {
					hasAdjacentSymbol = true
				}
			}
		}
	}
	return hasAdjacentSymbol
}

var symbolRegex = regexp.MustCompile(`[^.\d]`)

func checkIfSymbol(symbol byte) bool {
	return symbolRegex.MatchString(string(symbol))
}
