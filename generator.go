package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GeneratorExample() {
	fmt.Println("--------------------------")

	duration := 5 * time.Second

	doneChannel := make(chan bool)
	timeout := time.After(duration)

	randomFetcher := func() int { return rand.Intn(1000) }

	for {
		select {
		case <-timeout:
			close(doneChannel)
			fmt.Println("Timeout! Generator stopped.")
			return
		case randomNumber := <-repeatIt(doneChannel, randomFetcher):
			fmt.Println(randomNumber)
		}
	}

}

func repeatIt[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)

		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}
