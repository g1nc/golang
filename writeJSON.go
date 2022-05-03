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

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
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
	saveToJSON(os.Stdout, myRecord)
}
