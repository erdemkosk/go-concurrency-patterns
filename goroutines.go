package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution in Go. It's a concurrent unit of work that can be executed independently within a Go program.
//  Goroutines are managed by the Go runtime and allow multiple functions to run concurrently.
//  They're different from threads in other programming languages in that they're managed at a higher level by the Go runtime, enabling efficient scheduling and utilization of resources.
//  Goroutines are relatively cheap to create and are fundamental to building concurrent and scalable applications in Go.

func GoRoutinesExample() {
	fmt.Println("--------------------------")
	channel1 := make(chan string)
	channel2 := make(chan string)
	bufferedChannel := make(chan string, 2)

	go goRoutine1(channel1)
	go goRoutine2(channel2)

	go bufferedGoRoutineName(bufferedChannel)
	go bufferedGoRoutineSurname(bufferedChannel)

	msg1 := <-channel1
	msg2 := <-channel2
	name, surname := <-bufferedChannel, <-bufferedChannel

	fmt.Println(msg1)
	fmt.Println(msg2)

	fmt.Printf("Your name is %s %s\n", name, surname)

	fmt.Println("Hello from GoRoutines!")
}

func goRoutine1(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "1"
}

func goRoutine2(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "2"
}

func bufferedGoRoutineName(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "Erdem"
}

func bufferedGoRoutineSurname(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "Köşk"
	c <- "Out of buffered value"

	// so if we use for loop to send anything to channel after that to prevent deadlock it should use close(channel) and close the channel for sending new values
}
