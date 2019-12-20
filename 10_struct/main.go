package main

import "fmt"

func main(){
    type People struct{
        name string
        age int
        country string
        gender string
    }

    //var people People
    //people.name = "bob"
    //people.age = 21
    //people.country = "America"
    //people.gender = "m"

    //people := new(People)
    //people.name = "bob"
    //people.age = 21
    //people.country = "America"
    //people.gender = "m"

    //people := People{name:"bob", age:21, country:"America", gender:"m"}
    people := People{"bob", 21, "America", "m"}

    fmt.Println(people)
    fmt.Println("people.name:",people.name)
}
