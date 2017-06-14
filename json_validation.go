// [_Command-line arguments_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.
// For example, `go run hello.go` uses `run` and
// `hello.go` arguments to the `go` program.

package main

import "os"
import "fmt"
import "encoding/json"
import "github.com/asaskevich/govalidator"
import "net"
import "io/ioutil"
import "strconv"
import "text/template"
import "log"



type JsonData struct {

        Id int `valid:"-"`
        Client_id int `valid:"-"`
        Device_id int `valid:"-"`
        Domainname string `valid:"dns"`
        Ipaddress string `valid:"ip"`
        Domainport int `valid:"-"`
        Sitename string `valid:"-"`
        Status string `valid:"-"`
        Updated_at string `valid:"-"`
        EnableSsl int `valid:"-"`
        EnableGzip int `valid:"-"`
        SslKey string `valid:"-"`
        SslCrt string `valid:"-"`
        WcRedirect int `valid:"-"`

}

        var env JsonData
	var path string



func main() {

    // `os.Args` provides access to raw command-line
    // arguments. Note that the first value in this slice
    // is the path to the program, and `os.Args[1:]`
    // holds the arguments to the program.


	argsWithProg := os.Args[1]


var respBytes = []byte(argsWithProg)

//	var env JsonData
        if err := json.Unmarshal(respBytes, &env); err != nil {
                fmt.Println(err)
        }


	if ok, err := govalidator.ValidateStruct(env); err != nil {
        	panic(err)
        } else {
                fmt.Printf("OK: %v\n", ok)
        }

ip := net.ParseIP(env.Ipaddress)

if ip == nil{
                fmt.Println("Invalid IP")
                return
        }



	fmt.Println(env.Id)
        fmt.Println(env.Client_id)
        fmt.Println(env.Device_id)
        fmt.Println(env.Domainname)
        fmt.Println(env.Ipaddress)
        fmt.Println(env.Domainport)
        fmt.Println(env.Sitename)
        fmt.Println(env.Status)
        fmt.Println(env.Updated_at)
        fmt.Println(env.EnableSsl)
        fmt.Println(env.EnableGzip)
        fmt.Println(env.SslKey)
        fmt.Println(env.SslCrt)
        fmt.Println(env.WcRedirect)

       
//fmt.Printf("%+v\n", env)

//fmt.Println(respBytes)

	createCDN(&env)


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
                {env.Ipaddress, "", false},
//              {"", "", false},
//              {"", "", false},
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


func createCDN(env* JsonData){
//	fmt.Println(env.Ipaddress)
	path = strconv.Itoa(env.Device_id)+".conf"
	createFile()
	writeFile()
	readFile()
//	deleteFile()

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

func writeFile() {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// write some text to file
	_, err = file.WriteString("This is a test\n")
	if err != nil {
		fmt.Println(err.Error())
		return //must return here for defer statements to be called
	}
	_, err = file.WriteString("mari belajar golang\n")
	if err != nil {
		fmt.Println(err.Error())
		return //same as above
	}

	// save changes
	err = file.Sync()
	if err != nil {
		fmt.Println(err.Error())
		return //same as above
	}
}

func readFile() {
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// read file
	var text = make([]byte, 1024)
	n, err := file.Read(text)
	if n > 0 {
		fmt.Println(string(text))
	}
	//if there is an error while reading
	//just print however much was read if any
	//at return file will be closed
}

func deleteFile() {
	// delete file
	var err = os.Remove(path)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
