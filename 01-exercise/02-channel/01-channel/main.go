package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(a, b, d int) {
		c := a + b + d
		ch <- c
	}(1, 2, 7)
	// get the value computed from goroutine
	r := <-ch
	fmt.Printf("computed value %v\n", r)
}
