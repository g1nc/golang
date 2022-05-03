package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    var N = 40000000
    split := make([]map[int]int, 200)
    for i := range split {
        split[i] = make(map[int]int)
    }
    for i := 0; i < N; i++ {
        value := int(i)
        split[i%200][value] = value
    }
    start := time.Now()
    runtime.GC()
    fmt.Println("It took GC()", time.Since(start), "to finish.")
    _ = split[0][0]
}
