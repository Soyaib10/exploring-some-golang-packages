package main

import (
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	// 1 worker, max 2 tasks in the queue
	pool := pond.NewPool(1, pond.WithQueueSize(2))

	fmt.Println("📤 Submitting tasks to a pool with 1 worker and queue size 2...")

	// Task 1: Runs immediately on the worker
	pool.Submit(func() {
		fmt.Println("▶️ Task 1: Running immediately")
		time.Sleep(2 * time.Second) // Takes a long time
		fmt.Println("✅ Task 1: Done")
	})

	// Task 2 & 3: Sit in the queue
	pool.Submit(func() { fmt.Println("▶️ Task 2: Waiting in queue") })
	pool.Submit(func() { fmt.Println("▶️ Task 3: Waiting in queue") })

	// Task 4: Queue is full! TrySubmit returns (Task, bool)
	_, ok := pool.TrySubmit(func() { fmt.Println("▶️ Task 4: Trying to enter...") })
	if !ok {
		fmt.Println("❌ Task 4: REJECTED (Queue is full)")
	}

	_, ok2 := pool.TrySubmit(func() { fmt.Println("▶️ Task 5: Trying to enter...") })
	if !ok2 {
		fmt.Println("❌ Task 5: REJECTED (Queue is full)")
	}

	pool.StopAndWait()
	fmt.Println("🏁 All done.")
}
