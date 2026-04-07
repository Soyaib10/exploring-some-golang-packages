package main

import (
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	// Start with only 1 worker
	pool := pond.NewPool(1)

	fmt.Println("🚀 Pool started with 1 worker.")

	// Submit 6 slow tasks.
	// With 1 worker, each takes 1 second. Total time: 6 seconds.
	for i := 1; i <= 6; i++ {
		taskNum := i
		pool.Submit(func() {
			fmt.Printf("  ▶️ Task %d started\n", taskNum)
			time.Sleep(1 * time.Second)
			fmt.Printf("  ✅ Task %d done\n", taskNum)
		})
	}

	// Wait 2 seconds... (2 tasks will have finished by now)
	time.Sleep(2 * time.Second)

	fmt.Println("\n📈 Resizing pool to 3 workers...")
	pool.Resize(3)

	fmt.Println("⏳ Now 3 tasks run in parallel. The rest will finish faster!")

	// Wait for everything to finish
	pool.StopAndWait()

	fmt.Println("\n🏁 All tasks done. No restart needed!")
}
