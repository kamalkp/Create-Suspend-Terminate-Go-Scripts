// [_Command-line arguments_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.
// For example, `go run hello.go` uses `run` and
// `hello.go` arguments to the `go` program.

package main

import "os"
import "fmt"
import "encoding/json"
import "net"


        type JsonData struct {



        Id int64
        Client_id int64
        Device_id int64
        Domainname string
        Ipaddress string
        Domainport int64
        Sitename string
        Status string
        Updated_at string
        EnableSsl int64
        EnableGzip int64
        SslKey string
        SslCrt string
        WcRedirect int64

}

	var env JsonData


func main() {

    // `os.Args` provides access to raw command-line
    // arguments. Note that the first value in this slice
    // is the path to the program, and `os.Args[1:]`
    // holds the arguments to the program.


//	if(ParseIP(env.Ipaddress) == nil)
//	return

	argsWithProg := os.Args[1]
	fmt.Println(argsWithProg)


	var respBytes = []byte(argsWithProg)


//	var env JsonData
        if err := json.Unmarshal(respBytes, &env); err != nil {
                fmt.Println(err)
        }
        // for the love of Gopher DO NOT DO THIS


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


        
	fmt.Printf("%+v\n", env)

	fmt.Println(respBytes)





//	ip := net.ParseIP(env.Ipaddress)

	fmt.Println("Unreachable code")



}
