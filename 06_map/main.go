package main

import "fmt"

func main(){
    person := make(map[string]string)
    person["firstname"] = "obama"
    person["lastname"] = "balake"
    fmt.Println("person:", person)
    fmt.Println("person['firstname']:", person["firstname"])
}
