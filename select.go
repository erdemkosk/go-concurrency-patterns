package main

import (
	"fmt"
	"time"
)

func SelectExample() {
	fmt.Println("--------------------------")
	channel3 := make(chan string)
	channel4 := make(chan string)

	charChannel := make(chan string, 5)

	go goRoutine3(channel3)
	go goRoutine4(channel4)

	go fillBufferedChannel(charChannel)

	select {
	case msg1 := <-channel3:
		fmt.Println("Received from Channel 3:", msg1)
	case msg2 := <-channel4:
		fmt.Println("Received from Channel 4:", msg2)
	}

	var result string
	for i := 0; i < 5; i++ {
		select {
		case char := <-charChannel:
			result += char
		}
	}

	fmt.Println("Received:", result)
	fmt.Println("Hello From Select!")
}

func goRoutine3(c chan string) {
	time.Sleep(4 * time.Second)
	c <- "3"
}

func goRoutine4(c chan string) {
	time.Sleep(4 * time.Second)
	c <- "4"
}

func fillBufferedChannel(c chan string) {
	chars := []string{"e", "r", "d", "e", "m"}

	for _, v := range chars {
		c <- v
	}

	close(c)
}
