package main

import (
	"fmt"
	"time"
)

func longOperation(done chan<- bool) {
	fmt.Printf("Started long operation...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
	done <- true
}
func main() {
	done := make(chan bool)
	go longOperation(done)
	<-done
	fmt.Println("Back to main")
}
