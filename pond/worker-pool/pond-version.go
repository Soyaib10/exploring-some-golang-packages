package main

import (
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	// Create a pool with max 3 workers (riders)
	pool := pond.NewPool(3)

	// Submit 10 orders as tasks
	for i := 1; i <= 10; i++ {
		orderNum := i // capture loop variable

		pool.Submit(func() {
			fmt.Printf("🏍️ Rider processing Order #%d\n", orderNum)
			time.Sleep(1 * time.Second) // Simulate delivery
			fmt.Printf("✅ Order #%d delivered!\n\n", orderNum)
		})
	}

	// Stop pool and wait for ALL submitted tasks to complete
	pool.StopAndWait()

	fmt.Println("🎉 All orders delivered! Riders clocked out.")
}
