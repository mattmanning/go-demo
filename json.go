package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"firstname"`
	Age  int    `json:"age"`
}

func main() {
	byt := []byte(`{"firstname":"Matt","age":38}`)
	p := person{}
	_ = json.Unmarshal(byt, &p)
	fmt.Println(p)
}
