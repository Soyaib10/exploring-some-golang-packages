package main

import (
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	fmt.Println("🚀 Using the default global pool...")

	// Submit tasks directly to the package, no "pool" variable needed
	task1 := pond.Submit(func() {
		fmt.Println("  ▶️ Task 1 (Default Pool)")
		time.Sleep(1 * time.Second)
	})

	task2 := pond.Submit(func() {
		fmt.Println("  ▶️ Task 2 (Default Pool)")
		time.Sleep(1 * time.Second)
	})

	task3 := pond.Submit(func() {
		fmt.Println("  ▶️ Task 3 (Default Pool)")
		time.Sleep(1 * time.Second)
	})

	// Wait for the last task to finish
	// In a real scenario, you might want to use a Group to wait for all of them
	task1.Wait()
	task2.Wait()
	task3.Wait()

	fmt.Println("🏁 All tasks done!")
}
