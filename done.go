package main

import (
	"fmt"
	"time"
)

func DoneChannelExample() {
	fmt.Println("--------------------------")

	doneChannel := make(chan bool)

	go execDone(doneChannel)

	time.Sleep(3 * time.Second)

	close(doneChannel)

	fmt.Println("Hello from Done Channel!")

	// ıf we wont give any done channel it will work infinite time.

	// go func() {
	// 	for {
	// 		select {
	// 		default:
	// 			fmt.Println("Infinite Value")
	// 		}
	// 	}
	// }()

	// time.Sleep(1 * time.Hour)

}

func execDone(doneChannel <-chan bool) {
	for {
		select {
		case <-doneChannel:
			return
		default:
			fmt.Println("Infinite Value")
		}
	}
}
