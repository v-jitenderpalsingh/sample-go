package main

import (
    "fmt"

    "github.com/lib/pq"
)

func main() {
    fmt.Println("Hello, World!")
    // Reference pq so it's not removed by go mod tidy
    _ = pq.ErrNotSupported
}
