package main

import "fmt"

type People struct{
    name string
    age int
    country string
    gender string
}

func main(){

    var bob People
    bob.name = "bob"
    bob.age = 21
    bob.country = "America"
    bob.gender = "m"

    lili := new(People)
    lili.name = "lili"
    lili.age = 21
    lili.country = "America"
    lili.gender = "f"

    sam := People{name:"sam", age:21, country:"America", gender:"m"}
    jack := People{"jack", 21, "America", "m"}

    fmt.Println(bob)
    fmt.Println("bob.name:", bob.name)
    fmt.Println("sam.name:", sam.name)
    fmt.Println("jack.name:", jack.name)
}
