package main

import (
	"fmt"
	"time"
)

func client1(c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("Message from client1 ==> %v", i)
		time.Sleep(1500 * time.Millisecond)
	}
}

func client2(c chan string) {
	for i := 10; i < 15; i++ {
		c <- fmt.Sprintf("Message from client2 ==> %v", i)
		time.Sleep(3000 * time.Millisecond)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go client1(c1)
	go client2(c2)

	maxEmpty := 10
	currEmpty := 0
	for currEmpty <= maxEmpty {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-c1:
			fmt.Println("Received from client1: ", v)
		case v := <-c2:
			fmt.Println("Received from client2: ", v)
		// Without the default, we might have a deadlock !
		// This is due to the non-blocking behavior of select {} with a default case
		default:
			fmt.Println("No value received")
			currEmpty++
		}
	}
}
