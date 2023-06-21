package main

import (
	"fmt"
	"time"
)

func write(c chan string) {
	// names := []string{"Bob", "Alice", "Bobette", "John"} // 4th element can be written only when a slot is free in the channel!
	names := []string{"Bob", "Alice", "Bobette"}
	for _, n := range names {
		c <- n
		fmt.Printf("Wrote %v to channel (len=%v)\n", n, len(c))
	}
	close(c)
}

func main() {
	c := make(chan string, 3)
	fmt.Printf("Channel data: cap=%v, len=%v\n", cap(c), len(c))

	go write(c)
	time.Sleep(2 * time.Second)

	for v := range c {
		fmt.Printf("read value %v (len=%v)\n", v, len(c))
		time.Sleep(1 * time.Second)
	}
}
