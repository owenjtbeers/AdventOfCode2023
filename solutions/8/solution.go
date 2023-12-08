package sol8

import (
	mylib "aoc/lib"
	"fmt"
	"regexp"
)

type LRNode struct {
	name  string
	left  string
	right string
}

func Sol8pt1() int {
	scanner := mylib.ReadLines("./solutions/8/input.txt")

	index := 0
	var instruction string
	nodes := make(map[string]LRNode)
	for scanner.Scan() {
		line := scanner.Text()
		if index == 0 {
			instruction = line
		}
		node := parseNode(line)
		if node.name != "" {
			nodes[node.name] = node
		}
		index++
	}
	currentNode := nodes["AAA"]
	steps := 0
	for currentNode.name != "ZZZ" {
		instructionIndex := steps % len(instruction)
		instructionChar := string(instruction[instructionIndex])
		if instructionChar == "L" {
			currentNode = nodes[currentNode.left]
		} else if instructionChar == "R" {
			currentNode = nodes[currentNode.right]
		}
		// fmt.Println(currentNode.name)
		steps++
	}
	return steps
}

type GhostPathInfo struct {
	startingNode LRNode
	steps        int
	endingNode   LRNode
	currentNode  LRNode
}

func Sol8() int {
	scanner := mylib.ReadLines("./solutions/8/input.txt")

	index := 0
	var instruction string
	nodes := make(map[string]LRNode)
	for scanner.Scan() {
		line := scanner.Text()
		if index == 0 {
			instruction = line
		}
		node := parseNode(line)
		if node.name != "" {
			nodes[node.name] = node
		}
		index++
	}
	ghosts := []GhostPathInfo{}
	for key, value := range nodes {
		if string(key[2]) == "A" {
			ghostNode := GhostPathInfo{value, 0, LRNode{}, value}
			ghosts = append(ghosts, ghostNode)
		}
	}
	for index, ghost := range ghosts {
		steps := 0
		fmt.Println(string(ghost.currentNode.name[2]))
		for string(ghost.currentNode.name[2]) != "Z" {
			instructionIndex := steps % len(instruction)
			instructionChar := string(instruction[instructionIndex])
			if instructionChar == "L" {
				ghost.currentNode = nodes[ghost.currentNode.left]
			} else if instructionChar == "R" {
				ghost.currentNode = nodes[ghost.currentNode.right]
			}
			steps++
		}
		ghosts[index].endingNode = ghost.currentNode
		ghosts[index].steps = steps
	}
	fmt.Println(ghosts)
	pathMultiples := []int{}
	for _, ghost := range ghosts {
		pathMultiples = append(pathMultiples, ghost.steps)
	}
	fmt.Println(pathMultiples)
	return lcmOfArray(pathMultiples)
}

var pattern = `(\w+)\s*=\s*\((\w+),\s*(\w+)\)`
var re = regexp.MustCompile(pattern)

func parseNode(line string) LRNode {
	match := re.FindStringSubmatch(line)
	if len(match) < 4 {
		// The string did not match the pattern
		return LRNode{}
	}

	name := match[1]
	left := match[2]
	right := match[3]
	return LRNode{name, left, right}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
func lcmOfArray(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}
