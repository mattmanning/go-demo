package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	p := person{Name: "Matt", Age: 38}
	fmt.Println(p)
	fmt.Println(p.Age)
}
