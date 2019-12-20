package main

import "fmt"

func main(){
    a, b := 12, 35
    if a <= b{
        fmt.Printf("%d is less than %d\n", a, b) 
    }else{
        fmt.Printf("%d is less than %d\n", b, a)
    }
    
    sum := 0
    for i:=0;i<=100;i++{
        sum = sum + i
    }
    fmt.Println(sum)

    
}
