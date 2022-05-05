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

func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Provide a filename!")
		os.Exit(1)
	}

	filename := arguments[1]
	var myRecord Record
	err := loadFromJSON(filename, &myRecord)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(myRecord)
	}
}
