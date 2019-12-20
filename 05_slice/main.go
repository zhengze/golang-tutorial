package main

import "fmt"

func main(){
    var ages = [3]int{98, 43, 10}
    fmt.Println("ages:", ages)
    fmt.Println("ages[1:2]:", ages[1:2])
    fmt.Println("ages[1:]:", ages[1:])
    fmt.Println("ages[:]:", ages[:])
    fmt.Println("ages[0:0]:", ages[0:0])
}
