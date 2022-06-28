package main

import "fmt"

func function1(i int) int {
	return i + i
}

func function2(i int) int {
	return i * i
}

func funFunc(f func(int) int, v int) int {
	return f(v)
}

func main() {
	fmt.Println("function1:", funFunc(function1, 3))
	fmt.Println("function2:", funFunc(function2, 3))
	fmt.Println("Inline:", funFunc(func(i int) int { return i * i * i }, 3))
}
