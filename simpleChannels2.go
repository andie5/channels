package main

import (
	"fmt"
	"time"
)

func sum(x, y int, c chan int) {
	go func() {
		c <- x + y
	}()
}

func main3() {
	consumer := make(chan int)
	go sum(1, 2, consumer)
	time.Sleep(1e9)
	fmt.Printf("the result is: %v", <-consumer)
}
