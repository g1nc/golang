package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type logWritter struct {
	body string
}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// lw := logWritter{}
	// io.Copy(lw, resp.Body)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(string(b))
}

func (l logWritter) Write(bs []byte) (int, error) {
	// fmt.Println(string(bs))
	// fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
