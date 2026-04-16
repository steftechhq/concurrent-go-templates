package main

import (
	"fmt"
	"sync"
	"time"
)

//When we use a buffered channel, the sender goroutine will not block as long as there is space available in the buffer

//When we create a channel, we can specify the buffer capacity of a channel. Whenever a sender goroutine writes a
// message without any receiver consuming the message, the channel will store the message
//This means that while there is space in the buffer, our sender does not block.
//It will keep on storing messages as long as capacity remains in the buffer. Once the capacity is filled up, the sender will block again,

//This message buffer buildup can also happen if the receiving end is slow and does not consume fast enough to keep up with the rate of messages being produced.
//Once a receiver goroutine is available to consume the messages, the messages are fed to the receiver in the same order they were sent.
//  This happens even if there is no message sender goroutine
//Once the receiver goroutine consumes all the messages and the buffer empties, the receiver goroutine will again block.
// A receiver will block in cases when we don’t have a sender or when the sender is producing messages at a slower rate than the receiver can read them.

func receiver(messages chan int, wGroup *sync.WaitGroup) {
	msg := 0
	for msg != -1 {
		time.Sleep(1 * time.Second)
		msg = <-messages
		fmt.Println("Received:", msg)
	}
	wGroup.Done()
}

func main() {
	msgChannel := make(chan int, 3)
	wGroup := sync.WaitGroup{}
	wGroup.Add(1)
	go receiver(msgChannel, &wGroup)
	for i := 1; i <= 6; i++ {
		fmt.Println(time.Now().Format("15:04:05"), "Sending:", i)
		msgChannel <- i
	}
	msgChannel <- -1
	wGroup.Wait()
}
