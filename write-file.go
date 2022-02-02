package main

import (
	"fmt"
	"os"
)

func main() {
	bytes := []byte("Here is\nsome more text\n")

	file, err := os.Create("outfile.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	n, err := file.Write(bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if n != len(bytes) {
		fmt.Println("failed to write data")
		os.Exit(1)
	}
}
