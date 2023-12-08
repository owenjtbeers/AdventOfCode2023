package main

import (
	sol "aoc/solutions"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please provide a day to run.")
		return
	}

	command := args[1]

	switch command {
	case "1":
		fmt.Println(sol.Sol1())
	case "2":
		fmt.Println(sol.Sol2())
	case "3":
		fmt.Println(sol.Sol3())
	// case "4":
	// 	fmt.Println(sol.Sol4())
	// 	fmt.Println(sol.Sol4())
	// case "5":
	// 	fmt.Println(sol.Sol5())
	// case "6":
	// 	fmt.Println(sol.Sol5())
	case "7":
		fmt.Println(sol.Sol7())
	case "8":
		fmt.Println(sol.Sol8())
	default:
		fmt.Println("Unknown command:", command)
	}
}
