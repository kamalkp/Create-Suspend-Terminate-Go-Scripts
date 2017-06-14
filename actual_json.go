// [_Command-line arguments_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.
// For example, `go run hello.go` uses `run` and
// `hello.go` arguments to the `go` program.

package main

import "os"
import "fmt"
import "encoding/json"
import "net"
import "github.com/asaskevich/govalidator"
import "regexp"
//import "strconv"


// Basic regular expressions for validating strings
const (
	Email          string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	CreditCard     string = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
	ISBN10         string = "^(?:[0-9]{9}X|[0-9]{10})$"
	ISBN13         string = "^(?:[0-9]{13})$"
	UUID3          string = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	UUID4          string = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	UUID5          string = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	UUID           string = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	Alpha          string = "^[a-zA-Z]+$"
	Alphanumeric   string = "^[a-zA-Z0-9]+$"
	Numeric        string = "^[0-9]+$"
	Int            string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	Float          string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	Hexadecimal    string = "^[0-9a-fA-F]+$"
	Hexcolor       string = "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	RGBcolor       string = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	ASCII          string = "^[\x00-\x7F]+$"
	Multibyte      string = "[^\x00-\x7F]"
	FullWidth      string = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	HalfWidth      string = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	Base64         string = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	PrintableASCII string = "^[\x20-\x7E]+$"
	DataURI        string = "^data:.+\\/(.+);base64$"
	Latitude       string = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	Longitude      string = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	DNSName        string = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	IP             string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	URLSchema      string = `((ftp|tcp|udp|wss?|https?):\/\/)`
	URLUsername    string = `(\S+(:\S*)?@)`
	Hostname       string = ``
	URLPath        string = `((\/|\?|#)[^\s]*)`
	URLPort        string = `(:(\d{1,5}))`
	URLIP          string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
	URLSubdomain   string = `((www\.)|([a-zA-Z0-9]([-\.][-\._a-zA-Z0-9]+)*))`
	URL            string = `^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`
	SSN            string = `^\d{3}[- ]?\d{2}[- ]?\d{4}$`
	WinPath        string = `^[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
	UnixPath       string = `^(/[^/\x00]*)+/?$`
	Semver         string = "^v?(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"
	tagName        string = "valid"
	binary_range	string = "([01])"
)


var (
	rxEmail          = regexp.MustCompile(Email)
	rxCreditCard     = regexp.MustCompile(CreditCard)
	rxISBN10         = regexp.MustCompile(ISBN10)
	rxISBN13         = regexp.MustCompile(ISBN13)
	rxUUID3          = regexp.MustCompile(UUID3)
	rxUUID4          = regexp.MustCompile(UUID4)
	rxUUID5          = regexp.MustCompile(UUID5)
	rxUUID           = regexp.MustCompile(UUID)
	rxAlpha          = regexp.MustCompile(Alpha)
	rxAlphanumeric   = regexp.MustCompile(Alphanumeric)
	rxNumeric        = regexp.MustCompile(Numeric)
	rxInt            = regexp.MustCompile(Int)
	rxFloat          = regexp.MustCompile(Float)
	rxHexadecimal    = regexp.MustCompile(Hexadecimal)
	rxHexcolor       = regexp.MustCompile(Hexcolor)
	rxRGBcolor       = regexp.MustCompile(RGBcolor)
	rxASCII          = regexp.MustCompile(ASCII)
	rxPrintableASCII = regexp.MustCompile(PrintableASCII)
	rxMultibyte      = regexp.MustCompile(Multibyte)
	rxFullWidth      = regexp.MustCompile(FullWidth)
	rxHalfWidth      = regexp.MustCompile(HalfWidth)
	rxBase64         = regexp.MustCompile(Base64)
	rxDataURI        = regexp.MustCompile(DataURI)
	rxLatitude       = regexp.MustCompile(Latitude)
	rxLongitude      = regexp.MustCompile(Longitude)
	rxDNSName        = regexp.MustCompile(DNSName)
	rxURL            = regexp.MustCompile(URL)
	rxSSN            = regexp.MustCompile(SSN)
	rxWinPath        = regexp.MustCompile(WinPath)
	rxUnixPath       = regexp.MustCompile(UnixPath)
	rxSemver         = regex	p.MustCompile(Semver)
	binrange	 = regexp.MustCompile(binary_range)
)




        type JsonData struct {

        Id int `valid:"-"`
        Client_id int `valid:"-"`
        Device_id int `valid:"-"`
        Domainname string `valid:"dns"`
        Ipaddress string `valid:ip"`
        Domainport int `valid:"-"`
        Sitename string `valid:"url"`
        Status string `valid:"-"`
        Updated_at string `valid:"-"`
        EnableSsl int `valid:"-"`
        EnableGzip int `valid:"-"`
        SslKey string `valid:"UnixPath"`
        SslCrt string `valid:"UnixPath"`
        WcRedirect int `valid:"-"`

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

	id := govalidator.IsInt()
	

	
	if ok, err := govalidator.ValidateStruct(env); err != nil {
	panic(err)
	} else {
		fmt.Printf("OK: %v\n", ok)
	}



	if ip == nil{
                fmt.Println("Invalid IP")
                return
        }


//	fmt.Println("This is where we are : %t",email)
        



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
