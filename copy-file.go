package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("outfile.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := io.Copy(file, os.Stdin); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
