package main

import "fmt"

type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

func createStructure(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}
	return &myStructure{n, s, h}
}

func returnStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}
	return myStructure{n, s, h}
}

func main() {
	s1 := createStructure("Oleg", "Globin", 123)
	s2 := returnStructure("Oleg", "Globin", 123)

	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)

	fmt.Println(s1)
	fmt.Println(s2)
}
