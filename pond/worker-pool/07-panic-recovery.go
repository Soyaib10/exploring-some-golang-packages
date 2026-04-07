package main

import (
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	pool := pond.NewPool(3)

	// Task 1: Normal
	task1 := pool.SubmitErr(func() error {
		fmt.Println("🟢 Task 1: Running normally...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("✅ Task 1: Done")
		return nil
	})

	// Task 2: PANICS!
	task2 := pool.SubmitErr(func() error {
		fmt.Println("🟡 Task 2: About to crash...")
		time.Sleep(300 * time.Millisecond)
		panic("database connection lost!") // Pond catches this
	})

	// Task 3: Normal
	task3 := pool.SubmitErr(func() error {
		fmt.Println("🔵 Task 3: I'm still running despite the crash!")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("✅ Task 3: Done")
		return nil
	})

	// Check results
	fmt.Println("\n--- Checking Results ---")
	
	// Task 1 succeeded
	if err := task1.Wait(); err != nil {
		fmt.Printf("❌ Task 1 failed: %v\n", err)
	}

	// Task 2 panicked - pond returns it as an error!
	if err := task2.Wait(); err != nil {
		fmt.Printf("❌ Task 2 panicked: %v\n", err)
	}

	// Task 3 succeeded
	if err := task3.Wait(); err != nil {
		fmt.Printf("❌ Task 3 failed: %v\n", err)
	}

	pool.StopAndWait()

	fmt.Println("\n🏁 Program survived! No crash.")
}
