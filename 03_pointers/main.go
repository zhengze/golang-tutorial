package main

import "fmt"

func main(){
    var a int = 1
    b := &a
    fmt.Printf("a:%d\n", a)
    fmt.Printf("address:%p\n", b)
    fmt.Println(*b)
}
