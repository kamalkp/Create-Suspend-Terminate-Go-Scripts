package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	filename   = "b.txt"
//	start_data = "12345"
)

func printContents() {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("CONTENTS:", string(data))
}

func main() {

	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	newPos, err := f.Seek(-2, 2); if err != nil {
		panic(err)
	}

	if _, err := f.WriteAt([]byte("A"), newPos); err != nil {
		panic(err)
	}

	printContents()
}

