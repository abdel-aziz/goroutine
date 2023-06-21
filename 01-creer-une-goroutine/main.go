package main

import (
	"fmt"
	"time"
)

func LongOperation(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Printf("Operation finished ! Took %d s\n", duration)
}

func main() {
	fmt.Println("Starting first operation")
	go LongOperation(1)

	fmt.Println("Starting second operation")
	LongOperation(1)
}
