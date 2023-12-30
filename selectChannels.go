// Problem statement
// Create a random bit generator that is a program that produces a sequence of 100000 randomly generated 1’s and 0’s using a goroutine.

// Hint: You can use the select statement to randomly choose the cases for random generation

package main

import "fmt"

func main() {
	numbers := make(chan int)

	go func() {
		for {
			select {
			case <-numbers:
				fmt.Print("0 ")
			case <-numbers:
				fmt.Print("1 ")
			}
		}
	}()

	for i := 0; i < 100000; i++ {
		numbers <- i
	}

}
