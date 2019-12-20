package main

import "fmt"

func main(){
    ranks := make(map[string]int)
    ranks["america"] = 1
    ranks["china"] = 2
    for k, v := range ranks{
        fmt.Println(k, v)
    }

    delete (ranks, "america")
    fmt.Println(ranks)
}
