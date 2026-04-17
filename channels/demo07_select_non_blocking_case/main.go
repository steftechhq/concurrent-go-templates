package main

import (
	"fmt"
	"time"
)

//Another use case for using select is when we need to use channels in a non-blocking manner. Recall that when we were discussing mutexes,
// we saw that Go provides a non-blocking tryLock() operation. This function call tries to acquire the lock, but if the lock is being used,
// it will return immediately with a false return value. Can we adopt this pattern also for channel operations? For example, c
// an we try to read a message from a channel? However, if no messages are available, instead of blocking, we can have the current execution
// work on a default set of instructions

//The instructions under the default case will get executed if none of the other cases is available.
//This gives us the ability to try to access one or more channels, but if none is ready, we can do something else.

func sendMsgAfter(seconds time.Duration) <-chan string {
	messages := make(chan string)
	go func() {
		time.Sleep(seconds)
		messages <- "Hello"
	}()
	return messages
}

func main() {
	messages := sendMsgAfter(3 * time.Second)
	for {
		select {
		case msg := <-messages:
			fmt.Println("Message received:", msg)
			return
		default:
			fmt.Println("No messages waiting")
			time.Sleep(1 * time.Second)
		}
	}
}
