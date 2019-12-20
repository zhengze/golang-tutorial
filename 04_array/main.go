package main

import "fmt"

func main(){
    //declare
    var prices [2] int
    //define
    prices[0] = 20
    prices[1] = 30
    fmt.Println("prices:",prices)
    fmt.Println("price[0]:",prices[0])

    //var names [2]string = [2]string{"tom", "bob"}
    //var names = [2]string{"tom", "bob"}
    //var names = []string{"tom", "bob"}
    //var names = [...]string{"tom", "bob"}
    names := [...]string{"tom", "bob"}

    fmt.Println("names:",names)
    fmt.Println("names[1]:",names[1])
}
