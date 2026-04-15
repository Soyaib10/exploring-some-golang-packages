package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	for i := range jobs {
		defer wg.Done()
		fmt.Printf("worker %d sending email to user %d", id, i)
		time.Sleep(1 * time.Second)
		fmt.Printf("worker %d done for user %d\n", id, i)
	}
}

func main() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, &wg)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		jobs <- i
		fmt.Printf("jobs enqued for user %d\n", i)
	}

	close(jobs)
	wg.Wait()
	fmt.Println("all mails are sent")
 }