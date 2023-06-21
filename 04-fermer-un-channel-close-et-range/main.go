package main

import (
	"fmt"
	"time"
)

func reader(c chan string) {
	fmt.Println("Start read")
	for n := range c {
		fmt.Println(n)
	}
	fmt.Println("Stop read")
}

func main() {
	c := make(chan string)
	go reader(c)

	c <- "Bob"
	c <- "Alice"
	fmt.Println("Closing channel")
	close(c)
	// c <- "Bobette" // panic, send on a closed channel

	time.Sleep(5 * time.Second)
}
