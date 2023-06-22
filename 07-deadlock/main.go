package main

import (
	"fmt"
	"time"
)

func hello(c chan string) {
	c <- "Hello there!"
}

func main() {
	c := make(chan string)
	//(deadlock when *all* goroutines are asleep)
	go hello(c)

	c <- "Hello"     // Causes a deadlock because all goroutines are writing to the channel
	fmt.Println(<-c) // Fix the deadlock!
	time.Sleep(1 * time.Second)
}
