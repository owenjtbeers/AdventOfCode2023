package mylib

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filePath string) *bufio.Scanner {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	// defer file.Close()

	// Create a scanner to read the file line by line
	fmt.Println("Reading file:", filePath)
	scanner := bufio.NewScanner(file)

	return scanner
}
