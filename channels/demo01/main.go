package main

import (
	"fmt"
	"time"
)

func main() {
	msgChannel := make(chan string)

	go receiver(msgChannel)
	fmt.Println("Sending HELLO...")
	msgChannel <- "HELLO"
	fmt.Println("Sending THERE...")
	msgChannel <- "THERE"
	fmt.Println("Sending STOP...")
	msgChannel <- "STOP"
}

func receiver(messages chan string) {
	msg := ""
	for msg != "STOP" {
		msg = <-messages
		fmt.Println("Received: ", msg)
	}
	time.Sleep(time.Duration(5) * time.Second)
}
