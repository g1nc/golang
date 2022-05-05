package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func main() {
	myRecord := Record{
		Name:    "Oleg",
		Surname: "Globin",
		Tel: []Telephone{
			{Mobile: false, Number: "1234=5678"},
			{Mobile: true, Number: "1234-5678"},
		},
	}
	rec, err := json.Marshal(myRecord)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(rec))

	var unRec Record
	err = json.Unmarshal(rec, &unRec)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(unRec)
}
