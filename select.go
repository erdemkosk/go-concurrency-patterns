package main

import (
	"fmt"
	"time"
)

// The select statement is used to listen for data coming from multiple channels concurrently, but the crucial aspect is the concurrent (or asynchronous) nature.
// Especially when multiple tasks are being performed, using the select statement allows simultaneous access to data from channels and optimizes the processing time of the program.

// For instance, when waiting for data from the network or different threads of execution, it enables the main thread to continue doing other tasks while waiting for data from channels.
//  Instead of waiting idly until data arrives from any channel, select concurrently monitors multiple channels and acts accordingly when data arrives.

// The main thread waits until data is received from any of the channels being listened to with the select statement. However, it can also continue with other tasks,
//  execute other threads of execution, or perform different operations.

// This feature allows programs to monitor and manage various tasks concurrently, preventing the program from blocking in unexpected situations.
// The use of select is highly beneficial in scenarios involving asynchronous operations or listening for data from various sources.

func SelectExample() {
	fmt.Println("--------------------------")
	channel3 := make(chan string)
	channel4 := make(chan string)

	go goRoutine3(channel3)
	go goRoutine4(channel4)

	select {
	case msg1 := <-channel3:
		fmt.Println("Received from Channel 3:", msg1)
	case msg2 := <-channel4:
		fmt.Println("Received from Channel 4:", msg2)
	}

	fmt.Println("Hello From Select!")
}

func goRoutine3(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "3"
}

func goRoutine4(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "4"
}
