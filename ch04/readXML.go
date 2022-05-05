package main

import (
	"encoding/xml"
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

func loadFromXML(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	decodeXML := xml.NewDecoder(in)
	err = decodeXML.Decode(key)
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
	err := loadFromXML(filename, &myRecord)
	if err == nil {
		fmt.Println("XML:", myRecord)
	} else {
		fmt.Println(err)
	}
}
