package main

import "fmt"

func WorkerPool(workerId int, job <-chan int, result chan int) {
	for j := range job {
		fmt.Println(fmt.Sprintf("Worker %d , Get Job %d", workerId, j))
		result <- <-job + workerId
	}
}

func main() {

	job := make(chan int, 3)
	result := make(chan int, 3)

	for i := 1; i <= 3; i++ {
		go WorkerPool(i, job, result)
	}

	for n := 1; n <= 3; n++ {
		job <- n
	}

	close(job)

	for j := 1; j <= 3; j++ {
		fmt.Println("Result is :=", <-result)
	}

}
