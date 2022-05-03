package main

import (
    "fmt"
    "runtime"
    "time"
)

type dataSlice struct {
    i, j int
}

func main() {
    var N = 40000000
    var structure []dataSlice
    for i := 0; i < N; i++ {
        value := int(i)
        structure = append(structure, dataSlice{value, value})
    }
    start := time.Now()
    runtime.GC()
    fmt.Println("It took GC()", time.Since(start), "to finish.")
    _ = structure[0]
}
