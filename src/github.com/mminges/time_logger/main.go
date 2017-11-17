package main

import (
    "fmt"
    "time"
)

func main() {
    for {
        fmt.Println("the time is", time.Now())
        time.Sleep(30 * time.Second)
    }
}
