package main

import (
	"fmt"
	"sync"
	"time"
)

func ping(c chan string, wg *sync.WaitGroup) {
	for i := 1; i < 11; i++ {
		c <- fmt.Sprintf("ping %v", i)
	}
	wg.Done()
}

func pong(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 100; i < 600; i += 100 {
		c <- fmt.Sprintf("pong %v", i)
	}
}

func print(c chan string, wg *sync.WaitGroup) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}

}

func main() {
	var wg sync.WaitGroup
	c := make(chan string)
	wg.Add(2)

	go print(c, &wg)
	go ping(c, &wg)
	go pong(c, &wg)

	wg.Wait()
	//time.Sleep(10 * time.Second)
}
