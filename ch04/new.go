package main

import "fmt"

type aStruct struct {
	Text string
}

func main() {
	pS := new(aStruct)
	fmt.Println(pS)

	sP := new([]aStruct)
	fmt.Println(sP)

	*sP = make([]aStruct, 10)
	fmt.Println(sP)
}
