package main

import (
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	// ResultPool[string] means tasks will return a string
	pool := pond.NewResultPool[string](3)

	// Submit tasks that return a result
	task1 := pool.Submit(func() string {
		fmt.Println("🧮 Calculating delivery fee for Order #1 (KFC)...")
		time.Sleep(500 * time.Millisecond)
		return "Fee: ৳50"
	})

	task2 := pool.Submit(func() string {
		fmt.Println("🧮 Calculating delivery fee for Order #2 (Pizza Hut)...")
		time.Sleep(300 * time.Millisecond)
		return "Fee: ৳75"
	})

	task3 := pool.Submit(func() string {
		fmt.Println("🧮 Calculating delivery fee for Order #3 (BFC)...")
		time.Sleep(400 * time.Millisecond)
		return "Fee: ৳60"
	})

	// Wait for each task and get the result
	result1, _ := task1.Wait()
	result2, _ := task2.Wait()
	result3, _ := task3.Wait()

	fmt.Println("\n📋 Results:")
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	pool.StopAndWait()
}
