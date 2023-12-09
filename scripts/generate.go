package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please provide a day to run.")
		return
	}

	day := args[1]

	dir := "./solutions/" + day
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	err1 := os.WriteFile("./solutions/"+day+"/input.txt", []byte(""), 0644)
	err2 := os.WriteFile("./solutions/"+day+"/solution.go", []byte("package sol"+day+"\n\nfunc Sol"+day+"() int {\n\treturn 0\n}"), 0644)
	err3 := os.WriteFile("./solutions/"+day+"/example1.txt", []byte(""), 0644)

	if err1 != nil {
		fmt.Println("Error creating input file:", err1)
		return
	}
	if err2 != nil {
		fmt.Println("Error creating solution file:", err2)
		return
	}

	if err3 != nil {
		fmt.Println("Error creating example file:", err3)
		return
	}
}
