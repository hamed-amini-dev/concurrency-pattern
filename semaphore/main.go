package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var maxGoRoutines = 5

func main() {

	semaphore := make(chan struct{}, maxGoRoutines)

	var waitGroup sync.WaitGroup
	for i := 0; i < 20; i++ {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()
			defer func() {
				<-semaphore
				log.Println("release semaphore")
			}()

			semaphore <- struct{}{}
			log.Println("add semaphore for i:=", i)

			fmt.Printf("Running task %d\n", i)
			time.Sleep(time.Second * 20)
		}()
	}

	waitGroup.Wait()

}
