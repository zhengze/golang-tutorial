package main

import "fmt"

func product(c chan<- int, factory string) {
    for i := 0; i < 10; i++ {
        fmt.Println(factory, " send ", i)
        c <- i
    }
}

func consume(c <-chan int){
    for j := 0; j < 10; j++ {
        fmt.Println("get ", <-c)
    }
}
func main() {
    ch1 := make(chan int, 10)
    ch2 := make(chan int, 10)
    go product(ch1, "factory1")
    go product(ch2, "factory2")
    consume(ch1)
    consume(ch2)
}
