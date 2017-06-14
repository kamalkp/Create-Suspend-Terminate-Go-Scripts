package main


import (
  "github.com/asaskevich/govalidator"
)


func main(){

	type User struct {
	FirstName string
	LastName string
}

str := govalidator.ToString(&User{"John", "Juan"})
println(str)



isValid := govalidator.IsURL(`http://user@passexample.come`)


isValidIP := govalidator.IsIP(`127.0.0.256`)


println(isValid ," ", isValidIP)


}
