package main

import (
	"fmt"
)

func operation() {

	fmt.Println("Done!")

}
func main() {

	fmt.Println("Back to main")
	fmt.Println("Back to main1")

	defer operation()

	fmt.Println("Back to main3")
	fmt.Println("Back to main4")

}
