package main

import (
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	pool := pond.NewPool(3)

	// Create a group for a batch of tasks
	group := pool.NewGroup()

	// Submit 5 tasks to the group
	for i := 1; i <= 5; i++ {
		orderNum := i
		group.Submit(func() {
			fmt.Printf("🏍️ Delivering Order #%d\n", orderNum)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("✅ Order #%d done\n", orderNum)
		})
	}

	fmt.Println("⏳ Waiting for all 5 orders to complete...")

	// One call waits for EVERYTHING in the group
	err := group.Wait()

	if err != nil {
		fmt.Printf("❌ Some delivery failed: %v\n", err)
	} else {
		fmt.Println("🎉 All 5 orders delivered!")
	}

	pool.StopAndWait()
}
