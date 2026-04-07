package main

import (
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	// A small pool to make metrics obvious
	pool := pond.NewPool(2, pond.WithQueueSize(5))

	// Fill the queue
	for i := 1; i <= 5; i++ {
		pool.Submit(func() {
			time.Sleep(2 * time.Second) // Make them wait in queue
		})
	}

	// Check stats immediately
	fmt.Printf("👷 Running Workers: %d\n", pool.RunningWorkers())
	fmt.Printf("📥 Waiting Tasks (Queue): %d\n", pool.WaitingTasks())
	fmt.Printf("📤 Submitted Tasks: %d\n", pool.SubmittedTasks())
	fmt.Printf("🗑️ Dropped Tasks: %d\n", pool.DroppedTasks())

	pool.StopAndWait()

	fmt.Println("\n--- After Completion ---")
	fmt.Printf("✅ Successful: %d\n", pool.SuccessfulTasks())
	fmt.Printf("❌ Failed: %d\n", pool.FailedTasks())
}
