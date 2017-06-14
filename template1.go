package main

import (
	"log"
	"os"
	"text/template"
	"fmt"
	"io/ioutil"
)

	var path string

func main() {
	// Define a template

	letterByte, err := ioutil.ReadFile("template.conf") // just pass the file name
    	if err != nil {
        	fmt.Print(err)
   	 }

	letter := string(letterByte) // convert content to a 'string'


// prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "", false},
//		{"", "", false},
//		{"", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	path = "b.txt"

	createFile()
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
        checkError(err)
        defer file.Close()
	



	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute( file , r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}



func createFile() {
        // detect if file exists
        var _, err = os.Stat(path)

        // create file if not exists
        if os.IsNotExist(err) {
                var file, err = os.Create(path)
                checkError(err) //okay to call os.exit()
                defer file.Close()
        }
}




func checkError(err error) {
        if err != nil {
                fmt.Println(err.Error())
                os.Exit(0)
        }
}

