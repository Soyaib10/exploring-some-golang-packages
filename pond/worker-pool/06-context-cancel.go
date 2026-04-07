package main

import (
	"context"
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	pool := pond.NewPool(3)

	// Create a context that cancels after 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	// Create a group that is bound to this context
	group := pool.NewGroupContext(ctx)

	fmt.Println("📤 Submitting 6 tasks (each takes 1 second)...")

	// Submit 6 tasks. With 3 workers, this would normally take 2 rounds (2s).
	for i := 1; i <= 6; i++ {
		taskNum := i
		group.SubmitErr(func() error {
			fmt.Printf("  ▶️ Task %d started\n", taskNum)

			// Simulate work that checks for cancellation
			select {
			case <-time.After(1 * time.Second):
				fmt.Printf("  ✅ Task %d finished\n", taskNum)
				return nil
			case <-ctx.Done():
				fmt.Printf("  ❌ Task %d CANCELED (context timed out)\n", taskNum)
				return ctx.Err()
			}
		})
	}

	// Wait for all tasks — or for the context to timeout
	err := group.Wait()

	if err != nil {
		fmt.Printf("\n⚠️  group.Wait() returned: %v\n", err)
	} else {
		fmt.Println("\n🎉 All tasks completed!")
	}

	// Always call cancel() to release resources
	cancel()

	pool.StopAndWait()
}
