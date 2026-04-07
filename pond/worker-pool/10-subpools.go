package main

import (
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	// Main pool has max 4 workers total
	pool := pond.NewPool(4)

	// Express subpool can use up to 2 workers
	// (They share the parent's 4 workers)
	express := pool.NewSubpool(2)

	fmt.Println("🚀 Submitting 4 Regular tasks and 3 Express tasks...")

	// Submit 3 Express tasks
	for i := 1; i <= 3; i++ {
		taskNum := i
		express.Submit(func() {
			fmt.Printf("🚄 Express Task %d started\n", taskNum)
			time.Sleep(1 * time.Second)
			fmt.Printf("✅ Express Task %d done\n", taskNum)
		})
	}

	// Submit 4 Regular tasks
	for i := 1; i <= 4; i++ {
		taskNum := i
		pool.Submit(func() {
			fmt.Printf("📦 Regular Task %d started\n", taskNum)
			time.Sleep(1 * time.Second)
			fmt.Printf("✅ Regular Task %d done\n", taskNum)
		})
	}

	// Wait for everything
	pool.StopAndWait()

	fmt.Println("🏁 All tasks complete!")
}
