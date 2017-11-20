package main

import (
    "fmt"
    "time"
    "flag"
)

func main() {
    var period = flag.Int("p",30,"")
    flag.Parse()
    for {
        fmt.Println("the time is", time.Now())
        time.Sleep(time.Duration(*period) * time.Second)
    }
}
