package main

import (
	"fmt"
	"log"
)

func fibo1(n int) int {
	if n <= 0 {
		log.Fatal("error")
	}

	count := 3
	var a, b int = 1, 1
	if n <= 2 {
		return 1
	}
	for count <= n {
		count = count + 1
		a, b = b, a+b
	}
	return b
}

func fibo2(n int) int {
	if n <= 0 {
		log.Fatal("error")
	}
	if n <= 2 {
		return 1
	}
	return fibo2(n-1) + fibo2(n-2)
}

func main() {
	fmt.Println(fibo1(5))
}
