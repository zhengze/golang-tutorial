package main

import (
	"fmt"
	"time"
)

func loop1() {
	var count int
	for count < 100 {
		count++
		fmt.Println(count)
		time.Sleep(time.Second)
	}
}

func loop2() {
	var count int = 100
	for {
		count--
		fmt.Println(count)
		time.Sleep(time.Second)
	}
}

func main() {
	go loop1()
	loop2()
}
