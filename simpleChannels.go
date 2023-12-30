package main

import (
	"fmt"
	"time"
)

func main2() {

	firstCh := make(chan string)
	fmt.Println("sending")

	go testChannels(firstCh)
	// go func(){
	//     time.Sleep(5 * time.Second)
	//     fmt.Printf("received %v\n", <- firstCh)
	// }()
	firstCh <- "hello" // putting hello on the channel
	fmt.Println("sent")

}

func testChannels(receiver chan string) {
	time.Sleep(5 * time.Second)
	fmt.Printf("received %v\n", <-receiver)
}
