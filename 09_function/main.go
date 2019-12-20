package main

import "fmt"

func getage(name string) int{
    peoples := make(map[string]int)
    peoples["tom"] = 21
    peoples["rose"] = 22
    return peoples[name]
}

func main(){
    name := "rose"
    fmt.Printf("%s is %d years old.\n", name, getage("rose"))
}
