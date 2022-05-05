package main

import (
	"encoding/json"
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
		return
	}

	var myRecord Record
	filename := arguments[1]
	err := loadFromJSON(filename, &myRecord)
	if err == nil {
		fmt.Println("JSON:", myRecord)
	} else {
		fmt.Println(err)
	}

	myRecord.Name = "Dimitris"
	xmlData, _ := xml.MarshalIndent(myRecord, "", "  ")
	xmlData = []byte(xml.Header + string(xmlData))
	fmt.Println("XML:", string(xmlData))

	data := Record{}
	err = xml.Unmarshal(xmlData, &data)
	if err != nil {
		fmt.Println("Unmarshalling from XML", err)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling to JSON", err)
		return
	}

	_ = json.Unmarshal([]byte(result), &myRecord)
	fmt.Println("JSON:", myRecord)
}
