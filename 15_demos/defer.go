package main

import (
    "fmt"
)

func foo(){
    defer fmt.Println("foo")
}

func bar(){
    defer fmt.Println("bar")
}

func main(){

    fmt.Println("defer begin")
    // 将defer放入延迟调用栈
    defer fmt.Println(1)
    foo()
    defer fmt.Println(2)
    defer bar()
    // 最后一个放入, 位于栈顶, 最先调用
    defer fmt.Println(3)
    fmt.Println("defer end")
}
