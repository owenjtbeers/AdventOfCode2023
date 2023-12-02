package sol2

import (
	mylib "aoc/lib"
	"fmt"
	"strconv"
	"strings"
)

func Sol2() int {
	// Read in lines
	scanner := mylib.ReadLines("./solutions/2/input.txt")
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line here
		_, powerOfNumberOfCubes := parseInput(line)
		sum += powerOfNumberOfCubes
	}
	return sum
}

func parseInput(line string) (int, int) {
	gameSplit := strings.Split(line, ":")
	id := strings.Split(gameSplit[0], " ")[1]
	draws := strings.Split(gameSplit[1], ";")
	// isInputValid := true
	maxMap := make(map[string]int)
	for _, draw := range draws {
		drawSplit := strings.Split(draw, ",")
		drawMap := make(map[string]int)
		for _, info := range drawSplit {
			infoSplit := strings.Split(info, " ")
			color := infoSplit[2]
			value, err := strconv.Atoi(infoSplit[1])
			if err != nil {
				fmt.Println(err.Error())
			}
			drawMap[color] = value
		}
		for key, value := range drawMap {
			_, exists := maxMap[key]
			if exists && value > maxMap[key] {
				maxMap[key] = value
			} else if !exists {
				maxMap[key] = value
			}
		}
	}
	intId, _ := strconv.Atoi(id)
	powerOfNumberOfCubes := powerOfNumberOfCubes(maxMap)
	fmt.Println(intId, maxMap, powerOfNumberOfCubes)
	return intId, powerOfNumberOfCubes
}

func powerOfNumberOfCubes(numberOfCubes map[string]int) int {
	power := 0
	for _, value := range numberOfCubes {
		if power != 0 {
			power *= value
		} else {
			power = value
		}
	}
	return power
}

// var validGame = map[string]int{
// 	"red":   12,
// 	"green": 13,
// 	"blue":  14,
// }

// func verifyDraw(draw map[string]int) bool {
// 	isValid := true
// 	for key, value := range draw {
// 		if value > validGame[key] {
// 			isValid = false
// 		}
// 	}
// 	return isValid
// }
