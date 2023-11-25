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

	go goRoutine1(channel1)
	go goRoutine2(channel2)

	msg1 := <-channel1
	msg2 := <-channel2

	fmt.Println(msg1)
	fmt.Println(msg2)

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
