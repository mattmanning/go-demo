package main

import (
    "fmt"
    "time"
)

func f(msg string) {
    fmt.Println(msg)
}

func main() {

    f("direct")

    go f("goroutine")

    go func(msg string) {
        time.Sleep(time.Second)
        fmt.Println(msg)
    }("first")

    fmt.Println("second")

    time.Sleep(time.Second * 2)
}
