package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filename := "questions.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Content Comes in as byte array
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	//byte array gets converted to string at this point. And is also printed out.
	fmt.Println(string(content))
}
