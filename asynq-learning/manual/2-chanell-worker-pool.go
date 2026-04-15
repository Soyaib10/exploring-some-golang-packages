package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int) {
	for userID := range jobs {
		fmt.Printf("Worker %d sending email to user %d\n", id, userID)
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d done for user %d\n", id, userID)
	}
}

func main() {
	jobs := make(chan int, 100) // buffer size 100

	// ৩টা worker চালু করো
	for w := 1; w <= 3; w++ {
		go worker(w, jobs)
	}

	// ১০টা job দাও
	for i := 1; i <= 10; i++ {
		jobs <- i
		fmt.Printf("Job enqueued for user %d\n", i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Server crashed!")
	return
}
