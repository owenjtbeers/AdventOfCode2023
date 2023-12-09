package sol4

import (
	mylib "aoc/lib"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Sol4pt1() int {
	scanner := mylib.ReadLines("./solutions/4/input.txt")

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		winningNumbers, numbers := parseCard(line)
		fmt.Println(winningNumbers, numbers)
		winCount := 0
		for _, number := range winningNumbers {
			if contains(numbers, number) {
				winCount++
			}
		}
		sum += int(math.Pow(2, float64(winCount-1)))
		// for power := 0; power < winCount; power++ {
		// 	sum += int(math.Pow(2, float64(power)))
		// }
	}
	return sum
}

func Sol4() int {
	scanner := mylib.ReadLines("./solutions/4/input.txt")

	copiesMap := make(map[int]int)
	index := 0
	for scanner.Scan() {
		index++
		line := scanner.Text()
		winningNumbers, numbers := parseCard(line)
		fmt.Println(winningNumbers, numbers)
		winCount := 0
		for _, number := range winningNumbers {
			if contains(numbers, number) {
				winCount++
			}
		}
		copiesMap[index]++
		for count := 0; count < winCount; count++ {
			copiesMap[index+count+1] += copiesMap[index]
		}
	}
	sum := 0
	for _, value := range copiesMap {
		sum += value
	}
	fmt.Println(copiesMap)
	return sum
}

var digitRegex = regexp.MustCompile(`\d+`)

func parseCard(line string) ([]int, []int) {
	numbers := strings.Split(line, ":")[1]
	parsedNumbers := strings.Split(numbers, "|")
	winningNumbers, _ := convertStringsToInts(digitRegex.FindAllString(parsedNumbers[0], -1))
	cardNumbers, _ := convertStringsToInts(digitRegex.FindAllString(parsedNumbers[1], -1))

	return winningNumbers, cardNumbers
}

func convertStringsToInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, s := range strings {
		intVal, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints[i] = intVal
	}
	return ints, nil
}

func contains(array []int, number int) bool {
	for _, value := range array {
		if value == number {
			return true
		}
	}
	return false
}
