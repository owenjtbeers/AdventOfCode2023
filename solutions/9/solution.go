package sol9

import (
	mylib "aoc/lib"
	"fmt"
	"regexp"
	"strconv"
)

func Sol9pt1() int {
	scanner := mylib.ReadLines("./solutions/9/input.txt")

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := parseLine(line)
		nextNumber := nextNumber(numbers)
		// fmt.Println(nextNumber)
		sum += nextNumber
	}
	return sum
}

func Sol9() int {
	scanner := mylib.ReadLines("./solutions/9/input.txt")

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := parseLine(line)
		nextNumber := previousNumber(numbers)
		// fmt.Println(nextNumber)
		sum += nextNumber
	}
	return sum
}

var digitRegex = regexp.MustCompile(`-?\d+`)

func parseLine(line string) []int {
	stringNumbers := digitRegex.FindAllString(line, -1)
	numbers := []int{}
	for _, stringNumber := range stringNumbers {
		number, _ := strconv.Atoi(stringNumber)
		numbers = append(numbers, number)
	}
	return numbers
}

func computeDifferenceBetweenArrayItems(numbers []int) ([]int, bool) {
	differences := make([]int, len(numbers)-1)
	allZero := true
	for i := 0; i < len(numbers)-1; i++ {
		differences[i] = numbers[i+1] - numbers[i]
		if differences[i] != 0 {
			allZero = false
		}
	}
	return differences, allZero
}

func nextNumber(numbers []int) int {
	differences, allZeroes := computeDifferenceBetweenArrayItems(numbers)
	fmt.Println(numbers, differences)
	if allZeroes {
		return numbers[len(numbers)-1]
	}
	nextNumberOfDiff := nextNumber(differences)
	// fmt.Println(differences, nextNumberOfDiff)
	return numbers[len(numbers)-1] + nextNumberOfDiff
}

func previousNumber(numbers []int) int {
	differences, allZeroes := computeDifferenceBetweenArrayItems(numbers)
	fmt.Println(numbers, differences)
	if allZeroes {
		return numbers[0]
	}
	previousNumberOfDiff := previousNumber(differences)
	// fmt.Println(differences, nextNumberOfDiff)
	return numbers[0] - previousNumberOfDiff
}
