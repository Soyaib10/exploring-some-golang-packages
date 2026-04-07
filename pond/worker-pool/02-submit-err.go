package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	pool := pond.NewPool(3)

	// Submit a task that succeeds
	task1 := pool.SubmitErr(func() error {
		fmt.Println("📧 Sending welcome email...")
		time.Sleep(500 * time.Millisecond)
		return nil // Success
	})

	// Submit a task that fails
	task2 := pool.SubmitErr(func() error {
		fmt.Println("💳 Processing payment...")
		time.Sleep(300 * time.Millisecond)
		return errors.New("payment gateway timeout") // Error
	})

	// Submit another success
	task3 := pool.SubmitErr(func() error {
		fmt.Println("📊 Generating report...")
		time.Sleep(400 * time.Millisecond)
		return nil
	})

	// Wait for each task and check results
	err1 := task1.Wait()
	if err1 != nil {
		fmt.Printf("❌ Task 1 failed: %v\n", err1)
	} else {
		fmt.Println("✅ Task 1 succeeded")
	}

	err2 := task2.Wait()
	if err2 != nil {
		fmt.Printf("❌ Task 2 failed: %v\n", err2)
	} else {
		fmt.Println("✅ Task 2 succeeded")
	}

	err3 := task3.Wait()
	if err3 != nil {
		fmt.Printf("❌ Task 3 failed: %v\n", err3)
	} else {
		fmt.Println("✅ Task 3 succeeded")
	}

	pool.StopAndWait()
}
