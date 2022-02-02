package main

import "fmt"

func main() {
	s := []string{"a", "b"}
	fmt.Println(s)

	s = append(s, "c")
	fmt.Println(s)
}
