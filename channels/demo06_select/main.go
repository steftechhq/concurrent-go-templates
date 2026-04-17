package main

import (
	"fmt"
	"time"
)

// The select statement lets us group the read operations on multiple channels together,
//
//	blocking the goroutine until a message arrives on any of the channels .
//
// Once the message arrives on any of the channels, the goroutine is unblocked and the associate code for that channel read is run
// We can then decide what else to do—either continue with our execution or go back to wait for the next message by using the select statement again

//Channels are first-class objects, which means that we can store them as variables, pass or return them from functions, or even send them on a channel.

func main() {
	messagesFromA := writeEvery("Tick", 1*time.Second)
	messagesFromB := writeEvery("Tock", 3*time.Second)

	for {
		select {
		case msg1 := <-messagesFromA:
			fmt.Println(msg1)
		case msg2 := <-messagesFromB:
			fmt.Println(msg2)
		}
	}
}
func writeEvery(msg string, seconds time.Duration) <-chan string {
	messages := make(chan string)
	go func() {
		for {
			time.Sleep(seconds)
			messages <- msg
		}
	}()
	return messages
}
