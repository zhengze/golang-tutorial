package main

import "fmt"

func main(){
    // declare
    var name string
    // define
    name = "Bob"
    // declare and define
    var age int = 20
    // shorten define
    gender := "m"
    // const
    const Country = "America"

    firstname, lastname := "obama", "balake"

    fmt.Println(name, age, gender)
    fmt.Println(firstname, lastname)
    fmt.Println("Country:"+Country)
}
