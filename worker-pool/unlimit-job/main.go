package main

import (
	"fmt"
	"time"
)

func WorkerPool(workerId int, job <-chan int, result chan<- int) {
	for j := range job {
		fmt.Println(fmt.Sprintf("Worker %d , Get Job %d", workerId, j))
		result <- j
	}
}

func main() {

	job := make(chan int)
	result := make(chan int)

	for i := 1; i <= 3; i++ {
		go WorkerPool(i, job, result)
	}

	go func() {
		for i := 1; ; i++ {
			fmt.Println("Send Job :", i)
			job <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for j := 1; ; j++ {
		fmt.Println("Result is :=", <-result)
	}

}
