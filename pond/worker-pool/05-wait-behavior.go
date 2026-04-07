package main

import (
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	pool := pond.NewPool(2) // Only 2 workers
	group := pool.NewGroup()

	fmt.Println("📤 Submitting 4 tasks...")

	// Submit 4 tasks. Workers will start them immediately.
	for i := 1; i <= 4; i++ {
		taskNum := i
		group.Submit(func() {
			fmt.Printf("  ▶️ Task %d STARTED\n", taskNum)
			time.Sleep(1 * time.Second)
			fmt.Printf("  ✅ Task %d FINISHED\n", taskNum)
		})
	}

	fmt.Println("📌 All 4 tasks submitted. Now calling group.Wait()...")
	fmt.Println("⏳ group.Wait() blocks here until ALL finish...")

	// This blocks until ALL 4 finish
	group.Wait()

	fmt.Println("🎉 group.Wait() returned — all 4 tasks are done!")

	pool.StopAndWait()
}
