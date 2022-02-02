package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    f, err := os.Open("infile.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // read 1024 bytes at a time
    buf := make([]byte, 1024)

    for {
        n, err := f.Read(buf)
        if err == io.EOF {
            // there is no more data to read
            break
        }
        if err != nil {
            fmt.Println(err)
            continue
        }
        if n > 0 {
            fmt.Print(string(buf[:n]))
        }
    }
}
