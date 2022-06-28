package main

import "fmt"

func returnFunc() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}

func main() {
	i := returnFunc()
	j := returnFunc()

	fmt.Println("1:", i())
	fmt.Println("2:", i())
	fmt.Println("j1:", j())
	fmt.Println("j2:", j())
	fmt.Println("3:", i())
}
