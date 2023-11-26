package main

import (
	"fmt"
	"sort"
)

func PipelineExample() {
	fmt.Println("--------------------------")

	nums := []int{3, 5, 2, 6, 8, 0, 2, 1}

	stage1 := make(chan []int)
	go sortStage(nums, stage1)

	stage2 := make(chan int)
	go sliceStage(stage1, stage2)

	stage3 := make(chan int)
	go plusStage(stage2, stage3)

	<-stage3

	fmt.Println("Hello From Pipeline!")
}

func sortStage(nums []int, out chan<- []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	})
	out <- nums
	close(out)
}

func sliceStage(in <-chan []int, out chan<- int) {
	nums := <-in
	for _, v := range nums {
		out <- v
	}
	close(out)
}

func plusStage(in <-chan int, out chan<- int) {
	sum := 0
	for v := range in {
		sum += v
	}
	fmt.Println("Received:", sum)
	out <- sum
	close(out)
}
