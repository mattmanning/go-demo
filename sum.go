package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func main() {
	var sum int
	sum = plus(1, 2)

	fmt.Println(sum)
}
