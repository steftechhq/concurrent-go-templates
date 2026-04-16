package main

import (
	"fmt"
	"time"
)

// infinite channel receiver
// print zeroes when listening to the closed channel
func receiver(messages <-chan int) {
	for {
		msg := <-messages
		fmt.Println(time.Now().Format("15:04:05"), "Received:", msg)
		time.Sleep(1 * time.Second)
	}
}

func receiverThatStops(messages <-chan int) {
	for {
		msg, more := <-messages //get more flag; if the channel is still open
		fmt.Println(time.Now().Format("15:04:05"), "Received:", msg, more)
		time.Sleep(1 * time.Second)
		if !more {
			return
		}
	}
}

func receiverThatStops2(messages <-chan int) {
	for msg := range messages {
		fmt.Println(time.Now().Format("15:04:05"), "Received:", msg)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Receiver finished.")
}

func main() {
	testChannel()

}

func testChannel() {
	msgChannel := make(chan int)
	// go receiver(msgChannel)
	// go receiverThatStops(msgChannel)
	go receiverThatStops2(msgChannel)

	for i := 1; i <= 3; i++ {
		fmt.Println(time.Now().Format("15:04:05"), "Sending:", i)
		msgChannel <- i
		time.Sleep(1 * time.Second)
	}
	close(msgChannel)
	time.Sleep(3 * time.Second)
}
